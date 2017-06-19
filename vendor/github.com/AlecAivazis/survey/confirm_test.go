package survey

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/AlecAivazis/survey/core"
	"github.com/AlecAivazis/survey/terminal"
	"github.com/stretchr/testify/assert"
)

func init() {
	// disable color output for all prompts to simplify testing
	core.DisableColor = true
}

func TestConfirmRender(t *testing.T) {

	testError := fmt.Errorf("TEST")
	tests := []struct {
		title    string
		prompt   Confirm
		data     ConfirmTemplateData
		expected string
	}{
		{
			"Test Confirm question output with default true",
			Confirm{Message: "Is pizza your favorite food?", Default: true},
			ConfirmTemplateData{},
			`? Is pizza your favorite food? (Y/n) `,
		},
		{
			"Test Confirm question output with default false",
			Confirm{Message: "Is pizza your favorite food?", Default: false},
			ConfirmTemplateData{},
			`? Is pizza your favorite food? (y/N) `,
		},
		{
			"Test Confirm answer output",
			Confirm{Message: "Is pizza your favorite food?"},
			ConfirmTemplateData{Answer: "Yes"},
			"? Is pizza your favorite food? Yes\n",
		},
		{
			"Test Confirm error output",
			Confirm{Message: "Is pizza your favorite food?"},
			ConfirmTemplateData{Answer: "Yes", Error: &testError},
			`✘ Sorry, your reply was invalid: TEST
? Is pizza your favorite food? Yes
`,
		},
		{
			"Test Confirm with help but help message is hidden",
			Confirm{Message: "Is pizza your favorite food?", Help: "This is helpful"},
			ConfirmTemplateData{},
			"? Is pizza your favorite food? [? for help] (y/N) ",
		},
		{
			"Test Confirm help output with help message shown",
			Confirm{Message: "Is pizza your favorite food?", Help: "This is helpful"},
			ConfirmTemplateData{ShowHelp: true},
			`ⓘ This is helpful
? Is pizza your favorite food? (y/N) `,
		},
	}

	outputBuffer := bytes.NewBufferString("")
	terminal.Stdout = outputBuffer

	for _, test := range tests {
		outputBuffer.Reset()
		test.data.Confirm = test.prompt
		err := test.prompt.Render(
			ConfirmQuestionTemplate,
			test.data,
		)
		assert.Nil(t, err, test.title)
		assert.Equal(t, test.expected, outputBuffer.String(), test.title)
	}
}
