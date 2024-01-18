package tasks_test

import (
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"crynux_bridge/relay"
	"crynux_bridge/tasks"
	"crynux_bridge/tests"
	"image"
	"image/png"
	"io"
	"math/big"
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetTaskResult(t *testing.T) {

	err := tests.SyncToLatestBlock()
	assert.Equal(t, nil, err, "catchup error")

	sendTaskChan := make(chan int)
	go tasks.StartSendTaskOnChainWithTerminateChannel(sendTaskChan)

	getTaskCreationResultChan := make(chan int)
	go tasks.StartGetTaskCreationResultWithTerminateChannel(getTaskCreationResultChan)

	syncBlockChan := make(chan int)
	go tasks.StartSyncBlockWithTerminateChannel(syncBlockChan)

	addresses, privateKeys, err := tests.PrepareAccounts()
	assert.Nil(t, err, "error preparing accounts")

	err = tests.PrepareNetwork(addresses, privateKeys)
	assert.Nil(t, err, "error preparing the network")

	err = tests.PrepareTaskCreatorAccount(addresses[0], privateKeys[0])
	assert.Nil(t, err, "error preparing the task creator account")

	task, err := tests.NewTask()
	assert.Nil(t, err, "error creating task")

	time.Sleep(20 * time.Second)
	task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskBlockchainConfirmed)

	// Prepare the images
	// Calculate the pHash

	log.Debugln("calculating phash")

	numImages, err := models.GetTaskConfigNumImages(task.TaskArgs)
	assert.Nil(t, numImages, "error getting num_images in task args")

	var phash []byte
	images := make([]image.Image, numImages)
	pHashes := make([]string, numImages)

	for i := 0; i < numImages; i++ {

		img := tests.CreateImage()

		pHashBytes, err := blockchain.GetPHashForImage(img)
		assert.Nil(t, err, "error calculating phash for image")

		pHashes[i] = hexutil.Encode(pHashBytes)

		phash = append(phash, pHashBytes...)

		images[i] = img
	}

	log.Debugln("phash created: " + hexutil.Encode(phash))

	results := [3][]byte{
		phash,
		phash,
		phash,
	}

	err = tests.SubmitAndDiscloseResults(
		big.NewInt(int64(task.TaskId)),
		addresses,
		privateKeys,
		results)

	assert.Equal(t, nil, err, "error submitting result on chain")

	log.Debugln("task disclosed")

	time.Sleep(20 * time.Second)
	task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskPendingResult)

	log.Debugln("ready for upload results")

	imageReaders := make([]io.Reader, numImages)
	for i := 0; i < numImages; i++ {

		pr, pw := io.Pipe()
		currentImage := images[i]

		go func() {
			log.Debugln("encoding image in go routine...")
			err = png.Encode(pw, currentImage)
			assert.Nil(t, err, "error encoding png image")

			log.Debugln("encoding image completed")
			err = pw.Close()
			assert.Nil(t, err, "error closing png image writer")

			log.Debugln("encoding image pipe closed")
		}()

		imageReaders[i] = pr
	}

	log.Debugln("uploading task results")

	appConfig := config.GetConfig()
	appConfig.Blockchain.Account.Address = addresses[1]
	appConfig.Blockchain.Account.PrivateKey = privateKeys[1]

	log.Debugln("upload results using address: " + addresses[1])

	err = relay.UploadTaskResult(task.TaskId, imageReaders[:])
	assert.Nil(t, err, "error upload task results")

	log.Debugln("task results uploaded")

	// Prepare to download the results

	appConfig.Blockchain.Account.Address = addresses[0]
	appConfig.Blockchain.Account.PrivateKey = privateKeys[0]

	getResultChan := make(chan int)
	go tasks.StartDownloadResultsWithTerminateChannel(getResultChan)

	time.Sleep(20 * time.Second)

	// The results should be downloaded
	log.Debugln("download results using address: " + addresses[0])
	log.Debugln("task results should be downloaded already")

	task = tests.AssertTaskStatus(t, task.ID, models.InferenceTaskSuccess)

	taskResultFolder := path.Join(
		appConfig.DataDir.InferenceTasks,
		strconv.FormatUint(uint64(task.ID), 10),
	)

	for i := 0; i > numImages; i++ {
		downloadedImageFile := path.Join(taskResultFolder, strconv.Itoa(i)+".png")
		_, err := os.Stat(downloadedImageFile)
		assert.Nil(t, err, "image not downloaded")

		imageFile, err := os.Open(downloadedImageFile)
		assert.Nil(t, err, "error opening image file")

		img, err := png.Decode(imageFile)
		assert.Nil(t, err, "error decoding image")

		pHash, err := blockchain.GetPHashForImage(img)
		assert.Nil(t, err, "error calculating pHash for image")

		pHashStr := hexutil.Encode(pHash)

		assert.Equal(t, pHashes[i], pHashStr, "pHash for image mismatch")
	}

	t.Cleanup(func() {
		sendTaskChan <- 1
		getTaskCreationResultChan <- 1
		syncBlockChan <- 1
		getResultChan <- 1

		err = tests.ClearNetwork(addresses, privateKeys)
		assert.Equal(t, nil, err, "error clearing blockchain network")

		tests.ClearDB()
		err := tests.ClearDataFolders()
		assert.Equal(t, nil, err, "error clearing data folder")
	})
}
