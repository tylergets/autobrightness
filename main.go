package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	brightnessPath  = "/sys/class/backlight/intel_backlight/brightness"
	illuminancePath = "/sys/bus/iio/devices/iio:device0/in_illuminance_raw"
)

func main() {
	// Define command-line flags
	minBrightness := flag.Int("min", 1, "Minimum brightness value")
	sampleRate := flag.Duration("rate", time.Second, "Sample rate")
	flag.Parse()

	log.Printf("Starting automatic backlight adjustment (minimum brightness: %d, sample rate: %s)", *minBrightness, sampleRate)

	var lastBrightness int

	for {
		illuminance, err := readIlluminance()
		if err != nil {
			log.Println("Error reading illuminance:", err)
			continue
		}

		log.Println("Read illuminance:", illuminance)

		brightness, err := calculateBrightness(illuminance, *minBrightness)
		if err != nil {
			log.Println("Error calculating brightness:", err)
			continue
		}

		log.Println("Calculated brightness:", brightness)

		currentBrightness, err := getCurrentBrightness()
		if err != nil {
			log.Println("Error reading current brightness:", err)
			continue
		}

		if lastBrightness != 0 && lastBrightness != currentBrightness {
			log.Println("Brightness was changed externally. Exiting.")
			break
		}

		err = setBrightness(brightness)
		if err != nil {
			log.Println("Error setting brightness:", err)
			continue
		}

		log.Println("Set brightness:", brightness)
		lastBrightness = brightness

		time.Sleep(*sampleRate)
	}
}

func readIlluminance() (int, error) {
	data, err := ioutil.ReadFile(illuminancePath)
	if err != nil {
		return 0, err
	}

	strValue := strings.TrimSpace(string(data))
	rawValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}

	return rawValue, nil
}

func calculateBrightness(illuminance, minBrightness int) (int, error) {
	maxBrightness := 400

	// Map illuminance values from 0-27000 to brightness values from minBrightness-maxBrightness
	brightness := (illuminance * (maxBrightness - minBrightness) / 27000) + minBrightness

	// Make sure brightness is within the valid range of minBrightness-maxBrightness
	if brightness < minBrightness {
		brightness = minBrightness
	} else if brightness > maxBrightness {
		brightness = maxBrightness
	}

	return brightness, nil
}

func getCurrentBrightness() (int, error) {
	data, err := ioutil.ReadFile(brightnessPath)
	if err != nil {
		return 0, err
	}

	strValue := strings.TrimSpace(string(data))
	currentBrightness, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, err
	}

	return currentBrightness, nil
}

func setBrightness(brightness int) error {
	data := []byte(strconv.Itoa(brightness))

	return ioutil.WriteFile(brightnessPath, data, 0644)
}
