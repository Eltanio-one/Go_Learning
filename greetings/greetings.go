// define the name of our package
package greetings

// import fmt to enable text formatting
import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// return a greeting that embeds the name in a message
	// if no name given, return an error message
	if name == "" {
		return "", errors.New("No name provided")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages (map = dict)
	messages := make(map[string]string)
	// loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	// range = enumerate, returns index and value
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages.
// the return message is selected at random
func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Sapenin, %v",
		"Gwed %v ye chump",
		"Ay %v, ye breath's kickin",
	}

	// return a randomly selected message format by specifying a random index for
	// the slice of formats
	return formats[rand.Intn(len(formats))]
}
