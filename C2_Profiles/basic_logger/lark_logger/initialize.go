package lark_logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MythicMeta/MythicContainer/loggingstructs"
)

var larkWebhookURL = os.Getenv("WEBHOOK_DEFAULT_URL")

type LarkMessage struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func sendToLark(message string) {
	if larkWebhookURL == "" {
		return
	}

	larkMsg := LarkMessage{}
	larkMsg.MsgType = "text"
	larkMsg.Content.Text = message

	jsonData, err := json.Marshal(larkMsg)
	if err != nil {
		loggingstructs.AllLoggingData.Get("lark_logger").LogError(fmt.Errorf("sendToLark: failed to marshal"), "failed to marshal message", err)
		return
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(larkWebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		loggingstructs.AllLoggingData.Get("lark_logger").LogError(fmt.Errorf("sendToLark: failed to send"), "failed to send message", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		loggingstructs.AllLoggingData.Get("lark_logger").LogError(fmt.Errorf("sendToLark: lark returned status %d", resp.StatusCode), fmt.Sprintf("lark returned status %d", resp.StatusCode), nil)
	}
}

func Initialize() {
	myLoggerName := "lark_logger"
	myLogger := loggingstructs.LoggingDefinition{
		Name:           myLoggerName,
		Description:    "basic stdout debug logger with Lark webhook",
		LogToFilePath:  "mythic.log",
		LogLevel:       "debug",
		LogMaxSizeInMB: 20,
		LogMaxBackups:  10,
		NewCallbackFunction: func(input loggingstructs.NewCallbackLog) {
			msg := fmt.Sprintf("[Callback] %s - %v", input.Action, input)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input)
		},
		NewTaskFunction: func(input loggingstructs.NewTaskLog) {
			msg := fmt.Sprintf("[Task] %s - %v", input.Action, input.Data)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input.Data)
		},
		NewPayloadFunction: func(input loggingstructs.NewPayloadLog) {
			msg := fmt.Sprintf("[Payload] %s - %v", input.Action, input.Data)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input.Data)
		},
		NewKeylogFunction: func(input loggingstructs.NewKeylogLog) {
			msg := fmt.Sprintf("[Keylog] %s - %v", input.Action, input.Data)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input.Data)
		},
		NewCredentialFunction: func(input loggingstructs.NewCredentialLog) {
			msg := fmt.Sprintf("[Credential] %s - %v", input.Action, input.Data)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input.Data)
		},
		NewArtifactFunction: func(input loggingstructs.NewArtifactLog) {
			msg := fmt.Sprintf("[Artifact] %s - %v", input.Action, input.Data)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input.Data)
		},
		NewFileFunction: func(input loggingstructs.NewFileLog) {
			msg := fmt.Sprintf("[File] %s - %v", input.Action, input.Data)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input.Data)
		},
		NewResponseFunction: func(input loggingstructs.NewResponseLog) {
			msg := fmt.Sprintf("[Response] %s - %v", input.Action, input.Data)
			sendToLark(msg)
			loggingstructs.AllLoggingData.Get(myLoggerName).LogInfo(input.Action, "data", input.Data)
		},
	}
	loggingstructs.AllLoggingData.Get(myLoggerName).AddLoggingDefinition(myLogger)
}
