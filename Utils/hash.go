package utils

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
)

// the 4 necessary buffers for the MD5 algorithm
const (
	A = 0x67452301
	B = 0xEFCDAB89
	C = 0x98BADCFE
	D = 0x10325476
)

//padding function
func addPadding(str []byte) []byte { 
	// we need to store the original length of the URL/Message so we can append it after the padding
	sLen := len(str)
	str = append(str, 0x80)

	//add padding till there is only 64 bits left to be a multiple of 512
	//this is done by adding 0 as a padding after a single 1 bit
	//64 bits are left so we can store the original length of the message
	for len(str)%512 != 448 { 
		str = append(str, 0x00)
	}

	//here we add the message lenght to the padding
	bitLen := uint64(sLen) * 8
	byteLen := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteLen, bitLen)
	str = append(str, byteLen...)

	return str
}

func Hash(str []byte) []byte{
	str = addPadding(str)

	var buffers []uint32 = make([]uint32, 4)
		buffers[0] = A
		buffers[1] = B
		buffers[2] = C
		buffers[3] = D
	
	for i := 0; i < len(str); i+=64 {
		
		temp := str[i:i+64]
		var M [16]uint32
		for	j := 0; j < 16; j++ {
			M[j] = binary.LittleEndian.Uint32(temp[j*4 : j*4 + 4])
		}

		tA, tB, tC, tD := buffers[0], buffers[1], buffers[2], buffers[3]
		//then we perform 64 MD5 transformations
		//which won't be done here as there is already a pre-defined implementation by the crypto/md5 package

		//add the results of the 64 MD5 transformations

		buffers[0] += tA
		buffers[1] += tB
		buffers[2] += tC
		buffers[3] += tD
	}

	var hashed []byte = make([]byte, 16)
	binary.LittleEndian.PutUint32(hashed[0:4], buffers[0])
	binary.LittleEndian.PutUint32(hashed[4:8], buffers[1])
	binary.LittleEndian.PutUint32(hashed[8:12], buffers[2])
	binary.LittleEndian.PutUint32(hashed[12:16], buffers[3])

	return hashed
}

func MD5Hash(str string) string {
	//if we implemented the 64 rounds we can use the following function
	//return Hash(str)
	//else we are sticking with the pre-defined implementation by the crypto/md5 package
	hash := md5.Sum([]byte(str))
	hashed := hex.EncodeToString(hash[:])
	return hashed
}