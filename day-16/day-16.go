package main

import (
	"math"
	"strconv"
	. "utils"
)

func puzzle1(input string) int {
	binary := hexToBinary(input)
	return packet(&binary).versionNumberSum()
}

func puzzle2(input string) int {
	binary := hexToBinary(input)
	return packet(&binary).value
}

func packet(binary *string) Packet {
	version := binaryToDecimal(popBits(binary, 3))
	typeId := binaryToDecimal(popBits(binary, 3))

	if typeId == 4 {
		valueString := ""
		for foundZero := false; !foundZero; {
			foundZero = string((*binary)[0]) == "0"
			valueString += popBits(binary, 5)[1:5]
		}
		return Packet{version, []Packet{}, binaryToDecimal(valueString)}
	}

	lengthTypeId := binaryToDecimal(popBits(binary, 1))
	packets := []Packet{}

	if lengthTypeId == 0 {
		subPacketsLength := binaryToDecimal(popBits(binary, 15))

		subPacketsString := popBits(binary, subPacketsLength)
		for len(subPacketsString) > 0 {
			packets = append(packets, packet(&subPacketsString))
		}
	} else {
		packetCount := binaryToDecimal(popBits(binary, 11))

		for i := 0; i < packetCount; i++ {
			packets = append(packets, packet(binary))
		}
	}

	value := 0
	switch typeId {
	case 0:
		for _, packet := range packets {
			value += packet.value
		}
	case 1:
		value = 1
		for _, packet := range packets {
			value *= packet.value
		}
	case 2:
		value = math.MaxInt
		for _, packet := range packets {
			if packet.value < value {
				value = packet.value
			}
		}
	case 3:
		value = 0
		for _, packet := range packets {
			if packet.value > value {
				value = packet.value
			}
		}
	case 5:
		if packets[0].value > packets[1].value {
			value = 1
		} else {
			value = 0
		}
	case 6:
		if packets[0].value < packets[1].value {
			value = 1
		} else {
			value = 0
		}
	case 7:
		if packets[0].value == packets[1].value {
			value = 1
		} else {
			value = 0
		}
	}

	return Packet{version, packets, value}
}

type Packet struct {
	versionNumber int
	packets       []Packet
	value         int
}

func (packet Packet) versionNumberSum() (result int) {
	result += packet.versionNumber

	for _, packet := range packet.packets {
		result += packet.versionNumberSum()
	}

	return
}

func hexToBinary(h string) (b string) {
	for _, c := range h {
		b += map[string]string{
			"0": "0000",
			"1": "0001",
			"2": "0010",
			"3": "0011",
			"4": "0100",
			"5": "0101",
			"6": "0110",
			"7": "0111",
			"8": "1000",
			"9": "1001",
			"A": "1010",
			"B": "1011",
			"C": "1100",
			"D": "1101",
			"E": "1110",
			"F": "1111",
		}[string(c)]
	}

	return
}

func binaryToDecimal(binary string) int {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	Check(err)
	return int(decimal)
}

func popBits(s *string, n int) string {
	bits := (*s)[0:n]
	*s = (*s)[n:]
	return bits
}
