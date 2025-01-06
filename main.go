package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/url"
    "os"
    "strings"

    "github.com/d3tourrr/NomiKinGo"
    "github.com/joho/godotenv"
)

type Config struct {
    Topics []string `json:"topics"`
    Tones  []string `json:"tones"`
}

func GetPost(topic string, tone string) (string, error) {
    apiKey := os.Getenv("KIN_API_KEY")
    companionId := os.Getenv("KIN_COMPANION_ID")

    if apiKey == "" || companionId == "" {
        return "", fmt.Errorf("KIN_API_KEY and KIN_COMPANION_ID must be set in the environment")
    }

    kin := NomiKin.NomiKin{
        ApiKey: apiKey,
        CompanionId: companionId,
    }

    mess := "Generate a post about " + topic + " with a " + tone + " tone. No longer than 200 characters in length."
    post, err := kin.SendKindroidMessage(&mess)
    if err != nil {
        return "", err
    }

    return post, nil
}

func main() {
    if err := godotenv.Load(); err != nil {
        fmt.Printf("Error loading .env file: %v\n", err)
        return
    }

    file, err := os.Open("config.json")
    if err != nil {
        fmt.Printf("Error opening config file: %v\n", err)
        return
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := Config{}
    err = decoder.Decode(&config)
    if err != nil {
        fmt.Printf("Error decoding JSON: %v\n", err)
        return
    }

    topic := config.Topics[rand.Intn(len(config.Topics))]
    tone := config.Tones[rand.Intn(len(config.Tones))]

    post, err := GetPost(topic, tone)
    if err != nil {
        fmt.Printf("Error generating post: %v\n", err)
        return
    }

    encodedPost := url.QueryEscape(post)

    fmt.Println("Generated Post:")
    fmt.Println(post)
    fmt.Println("\nPost to Social Media:")
    fmt.Printf("   X: %s?text=%s\n", "https://twitter.com/intent/tweet", encodedPost)
    fmt.Printf("Bsky: %s?text=%s\n", "https://bsky.app/intent/compose", encodedPost)

    var runAgain string
    fmt.Println("\nGo again? (y/n)")
    fmt.Scanln(&runAgain)
    if strings.ToLower(runAgain) == "y" {
        main()
    }
}

