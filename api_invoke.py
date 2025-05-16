import json
import time
from datetime import datetime
from typing import Any, Dict, Optional, Tuple

import httpx
from eth_account import Account
from eth_account.signers.local import LocalAccount
from web3 import Web3


class Signer(object):
    def __init__(self, privkey: str) -> None:
        self.account: LocalAccount = Account.from_key(privkey)

    def sign(
        self, input: Dict[str, Any], timestamp: Optional[int] = None
    ) -> Tuple[int, str]:
        input_bytes = json.dumps(
            input, ensure_ascii=False, separators=(",", ":"), sort_keys=True
        ).encode("utf-8")
        if timestamp is None:
            timestamp = int(time.time())
        t_bytes = str(timestamp).encode("utf-8")

        data_hash = Web3.keccak(input_bytes + t_bytes)

        res = bytearray(self.account.signHash(data_hash).signature)
        res[-1] -= 27
        return timestamp, "0x" + res.hex()


client = httpx.Client(
    base_url="http://localhost:5029",
    timeout=180,
    verify=False,
)


root_privkey = "0x232b4536ed7f23d0fd34129ee49c2a47c52ffe2b85ce052a993283e79a87b2bd"
root_signer = Signer(root_privkey)


create_api_key_input = {}
timestamp, signature = root_signer.sign(create_api_key_input)

resp = client.post(
    f"/v1/api_key",
    json={"timestamp": timestamp, "signature": signature},
)
resp.raise_for_status()
create_api_key_output = resp.json()["data"]
api_key = create_api_key_output["api_key"]
expires_at = create_api_key_output["expires_at"]
print(f"API key: {api_key}")
print(f"Expires at: {datetime.fromtimestamp(expires_at)}")


role_input = {
    "api_key": api_key,
    "role": "chat",
}
timestamp, signature = root_signer.sign(role_input)

resp = client.post(
    f"/v1/api_key/{api_key}/role",
    json={"timestamp": timestamp, "signature": signature, "role": "chat"},
)
resp.raise_for_status()

role_input = {
    "api_key": api_key,
    "role": "image",
}
timestamp, signature = root_signer.sign(role_input)

resp = client.post(
    f"/v1/api_key/{api_key}/role",
    json={"timestamp": timestamp, "signature": signature, "role": "image"},
)
resp.raise_for_status()

use_limit = 1000
change_use_limit_input = {
    "api_key": api_key,
    "use_limit": use_limit,
}
timestamp, signature = root_signer.sign(change_use_limit_input)
resp = client.post(
    f"/v1/api_key/{api_key}/use_limit",
    json={"timestamp": timestamp, "signature": signature, "use_limit": use_limit},
)
resp.raise_for_status()

rate_limit = 6
change_rate_limit_input = {
    "api_key": api_key,
    "rate_limit": rate_limit,
}
timestamp, signature = root_signer.sign(change_rate_limit_input)
resp = client.post(
    f"/v1/api_key/{api_key}/rate_limit",
    json={"timestamp": timestamp, "signature": signature, "rate_limit": rate_limit},
)
resp.raise_for_status()

# delete_api_key_input: Dict[str, Any] = {
#     "api_key": api_key,
# }
# timestamp, signature = root_signer.sign(delete_api_key_input)
# resp = client.delete(
#     f"/v1/api_key/{api_key}",
#     params={"timestamp": timestamp, "signature": signature},
# )
# resp.raise_for_status()


print("Done")
