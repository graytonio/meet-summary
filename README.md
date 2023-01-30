# Transcript Summarizer

A tool for turing transcripts into a short readable summary using OpenAI and GPT3

## Usage
```
Usage:
  zoltan-bot [flags]

Flags:
  -h, --help                help for zoltan-bot
  -t, --transcript string   Transcript file for zoom meeting (default "./transcript.txt")
```

## Prereq

The script expects there to be a gpt.env file with a key called OPENAI_API_KEY with an OpenAI API key to query with. Or a shell variable with the same name and data.
