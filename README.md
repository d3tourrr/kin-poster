# kin-poster

This is a silly little project to generate a post that can be shared to X and Bsky by clicking a link. It's extremely simple and will probably never get another update.

## Usage

`.env` file should contain:

```
KIN_API_KEY=your_kin_api_key
KIN_COMPANION=your_kin_companion_id
```

`config.json` file should contain:
```json
{
    "topics": [
        "Cryptocurrency skepticism",
        "Sports commentary"
    ],
    "tones": [
        "Informative",
        "Sarcastic"
    ]
}
```

Add as many topics and tones as you want. Each run will randomly select a topic and tone to generate a post. The Kindroid you specify in your `.env` file will be sent a message to create a post based on the topic and tone, and then links to post that text will be generated.
