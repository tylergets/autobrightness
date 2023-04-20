README

This Go program automatically adjusts the brightness of the backlight of an Intel-based device based on the illuminance value read from the device's ambient light sensor. The program periodically reads the current illuminance value and calculates the appropriate brightness value to set the backlight to. The brightness value is then written to the device's brightness file, located at "/sys/class/backlight/intel_backlight/brightness".

This program takes three command-line arguments:

"min" (default: 1): the minimum brightness value to set the backlight to
"rate" (default: 1 second): the time interval between each illuminance reading and brightness adjustment
"exit-on-change" (default: false): if set to true, the program will exit if the brightness is changed externally (e.g. through a keyboard shortcut or system settings).
The program uses the following functions to implement its functionality:

readIlluminance(): reads the current illuminance value from the device's ambient light sensor.
calculateBrightness(illuminance int, minBrightness int) (int, error): calculates the appropriate brightness value to set the backlight to based on the current illuminance value and the minimum brightness value specified by the user.
getCurrentBrightness(): reads the current brightness value of the device's backlight.
setBrightness(brightness int) error: writes the specified brightness value to the device's brightness file.
To run the program, navigate to the directory containing the program file and run the following command:

go
Copy code
go run main.go
To specify command-line arguments, include them after the command:

arduino
Copy code
go run main.go -min=2 -rate=2s -exit-on-change=true
Note: This program is specifically designed for Intel-based devices with an ambient light sensor located at "/sys/bus/iio/devices/iio:device0/in_illuminance_raw" and a backlight file located at "/sys/class/backlight/intel_backlight/brightness". It may not work on devices with different file paths or hardware configurations.