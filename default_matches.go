package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var default_matches = map[string]string{
	// add some from slack
	"</3":  "broken_heart",
	"8)":   "sunglasses",
	"D:":   "anguished",
	":o)":  "monkey_face",
	":-*":  "kiss",
	":*":   "kiss",
	":×":   "kiss",
	"=)":   "smiley",
	"=-)":  "smiley",
	":-D":  "smile",
	";‑)":  "wink",
	":>":   "laughing",
	":->":  "laughing",
	">:(":  "angry",
	">:-(": "angry",
	"(:":   "slightly_smiling_face",
	"):":   "slightly_frowning_face",
	":\\":  "confused",
	":-\\": "confused",
	":‑p":  "stuck_out_tongue",
	":p":   "stuck_out_tongue",
	":b":   "stuck_out_tongue",
	":-b":  "stuck_out_tongue",
	";‑p":  "stuck_out_tongue_closed_eyes",
	";p":   "stuck_out_tongue_closed_eyes",
	";b":   "stuck_out_tongue_closed_eyes",
	";-b":  "stuck_out_tongue_closed_eyes",

	// Add some from https://en.wikipedia.org/wiki/List_of_emoticons
	// Smiley or happy face
	":‑)": "smiley",
	":)":  "smiley", // mattermost
	":-]": "smirk",  // mattermost
	":]":  "smirk",  // mattermost
	":-3": "smiley",
	":3":  "smiley",
	"8-)": "sunglasses",
	":-}": "smiley",
	":}":  "smiley",
	":c)": "smiley",
	":^)": "smiley",
	"=]":  "smiley",
	// Laughing, big grin, laugh with glasses, or wide-eyed surprise
	":‑D": "smile", // mattermost
	":D":  "smile", // mattermost
	"8‑D": "smile",
	"8D":  "smile",
	"x‑D": "laughing",
	"xD":  "laughing",
	"X‑D": "laughing",
	"XD":  "laughing",
	"=D":  "smile",
	"=3":  "smile",
	"B^D": "smile",
	// Very happy or double chin
	":-))": "grin",
	// Frown, sad, angry, pouting
	":‑(": "slightly_frowning_face",
	":(":  "slightly_frowning_face", // mattermost
	":‑c": "frowning_face",
	":c":  "frowning_face",
	":‑<": "frowning",
	":<":  "frowning",
	":[":  "rage", // mattermost
	":‑[": "rage",
	":@":  "rage",
	// Crying
	":'‑(": "cry",
	":'(":  "cry", //mattermost
	// Tears of happiness
	":'‑)": "joy",
	":')":  "joy",
	// Surprise, shock, yawn
	":‑O": "open_mouth",
	":O":  "open_mouth",
	":‑o": "open_mouth",
	":o":  "open_mouth",
	":-0": "open_mouth",
	"8‑0": "open_mouth",
	// wink, smirk
	";)":  "wink", //mattermost
	"*-)": "wink",
	"*)":  "wink",
	";‑]": "wink",
	";]":  "wink",
	// Tongue sticking out, cheeky/playful
	":‑P": "stuck_out_tongue",
	":P":  "stuck_out_tongue", //mattermost
	"X‑P": "stuck_out_tongue_closed_eyes",
	"XP":  "stuck_out_tongue_closed_eyes",
	"x‑p": "stuck_out_tongue_closed_eyes",
	"xp":  "stuck_out_tongue_closed_eyes",
	":‑Þ": "stuck_out_tongue",
	":Þ":  "stuck_out_tongue",
	":‑þ": "stuck_out_tongue",
	":þ":  "stuck_out_tongue",
	":‑b": "stuck_out_tongue",
	"d:":  "stuck_out_tongue",
	"=p":  "stuck_out_tongue",
	// Skeptical, annoyed, undecided, uneasy, hesitant
	":‑/":  "confused",
	":/":   "confused", // mattermost
	":‑.":  "confused",
	">:\\": "confused",
	"=/":   "confused",
	"=\\":  "confused",
	":L":   "confused",
	"=L":   "confused",
	":S":   "confounded", // mattermost
	// Straight face, no expression, indecision
	":‑|": "expressionless",
	":|":  "neutral_face", //mattermost
	// Embarassed, blushing
	":$": "flushed", // mattermost
	// Sealed lips or wearing braces, tongue-tied
	":‑X": "zipper_mouth_face",
	":X":  "zipper_mouth_face",
	":‑#": "zipper_mouth_face",
	":#":  "zipper_mouth_face",
	":‑&": "zipper_mouth_face",
	":&":  "zipper_mouth_face",
	// angel
	"O:‑)": "angel",
	"O:)":  "angel",
	"0:‑3": "angel",
	"0:3":  "angel",
	"0:‑)": "angel",
	"0:)":  "angel",
	"0;^)": "angel",
	// imp, devilish
	">:‑)": "smiling_imp",
	">:)":  "smiling_imp",
	"}:‑)": "smiling_imp",
	"}:)":  "smiling_imp",
	"3:‑)": "smiling_imp",
	"3:)":  "smiling_imp",
	">;)":  "smiling_imp",
}

func printDefaultMatches() {
	b := new(bytes.Buffer)
	e := json.NewEncoder(b)
	e.Encode(default_matches)
	fmt.Println(string(b.Bytes()))
}
