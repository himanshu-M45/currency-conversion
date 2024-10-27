package helper

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConversionRates_Success(t *testing.T) {
	// Use the actual XML file in the test
	rates, err := LoadConversionRates("conversion_rates.xml")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(rates))
	assert.Equal(t, 84.05, rates["INR"])
	assert.Equal(t, 1.0, rates["USD"])
	assert.Equal(t, 0.92, rates["EUR"])
}

func TestLoadConversionRates_FileNotFound(t *testing.T) {
	_, err := LoadConversionRates("non_existent_file.xml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open XML file")
}

func TestLoadConversionRates_InvalidXML(t *testing.T) {
	// Create a temporary XML file with invalid content
	invalidXMLContent := `<conversionRates><currency><code>USD</code><baseValue>1.0</baseValue></currency>`

	tmpFile, err := ioutil.TempFile("", "conversion_rates_*.xml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(invalidXMLContent))
	assert.NoError(t, err)
	tmpFile.Close()

	// Use the temporary file in the test
	_, err = LoadConversionRates(tmpFile.Name())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal XML")
}
