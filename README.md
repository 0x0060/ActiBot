
<p align="center">
  <img src="https://github.com/user-attachments/assets/e1122a69-44c7-41b5-bfad-73db67833144" alt="Project Logo" />
</p>


<p align="center">
  <a href="https://github.com/0x0060/ActiBot/issues"><img src="https://img.shields.io/github/issues/0x0060/ActiBot.svg?style=flat-square" alt="Issues" /></a>
  <a href="https://github.com/0x0060/ActiBot/network/members"><img src="https://img.shields.io/github/forks/0x0060/ActiBot.svg?style=flat-square" alt="Forks" /></a>
  <a href="https://github.com/0x0060/ActiBot/stargazers"><img src="https://img.shields.io/github/stars/0x0060/ActiBot.svg?style=flat-square" alt="Stars" /></a>
  <a href="https://github.com/0x0060/ActiBot/graphs/contributors"><img src="https://img.shields.io/github/contributors/0x0060/ActiBot.svg?style=flat-square" alt="Contributors" /></a>
  <a href="https://github.com/0x0060/ActiBot/blob/main/LICENSE"><img src="https://img.shields.io/github/license/0x0060/ActiBot.svg?style=flat-square" alt="License" /></a>
</p>

## Overview

The **ActiBot** is a powerful tool that monitors user activities within a Discord server, allowing users to share and display their currently playing music status from Spotify in their Discord profiles. Additionally, it provides an API that allows users to access activity data programmatically.

## Features

- 🎵 **Real-Time Music Status**: Automatically updates your Discord status to reflect what you're currently listening to on Spotify.
- 📊 **User Activity Tracking**: Monitors and tracks user activities in the Discord server.
- 📡 **API Access**: Provides an API for accessing user activity data.
- ⚙️ **Customizable Settings**: Easily configure the bot to suit your preferences.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Installation](#installation)
3. [Configuration](#configuration)
4. [Running the Bot](#running-the-bot)
5. [Usage](#usage)
6. [Contributing](#contributing)
7. [Contact](#contact)
8. [License](#license)



## Getting Started

### Prerequisites

To get started, ensure you have the following installed:

- **Go** (1.18 or later)
- **Discord Bot Token**: You can create a bot in the Discord Developer Portal.


## Installation

1. **Clone the repository**:
```bash
git clone https://github.com/0x0060/ActiBot.git
cd ActiBot
```

2. **Install dependencies**:
```bash
go mod tidy
```

3. **Configure your bot**:
   Create a `config.json` file in the root directory with the following structure:

```json
{
   "DiscordToken": "YOUR_DISCORD_BOT_TOKEN",
   "APIPort": ":8080"
}
```

   - Replace `YOUR_DISCORD_BOT_TOKEN` with your actual Discord bot token.
   - You can change the `APIPort` if you want the bot to run on a different port.
  

## Running the Bot

Once you have everything set up, you can run the bot with the following command:

```bash
go run cmd/bot/main.go
```

This will start the bot and it will begin listening for user activities on Spotify.


## Usage

Once the bot is running, it will automatically update your Discord status based on your current Spotify activity. Additionally, you can access user activity data via the provided API endpoints. Make sure to invite the bot to your Discord server and grant it the necessary permissions to manage your presence.

### Inviting the Bot to Your Server

To invite the bot to your Discord server, you need to use the following URL format:

https://discord.com/api/oauth2/authorize?client_id=YOUR_CLIENT_ID&permissions=YOUR_PERMISSIONS&scope=bot


- Replace `YOUR_CLIENT_ID` with your Discord bot's client ID.
- Set the desired permissions in `YOUR_PERMISSIONS` (e.g., `8192` for presence updates).


## Configuration

Alternatively, you can use a `config.yaml` file instead of `config.json`. Here’s how it should look:

```yaml
DiscordToken: "YOUR_DISCORD_BOT_TOKEN"
APIPort: ":8080"
```


## Contributing

We welcome contributions from the community! To contribute to this project, please follow these steps:

1. **Fork the repository**.
2. **Create a new branch** for your feature or bug fix.
3. **Make your changes** and commit them.
4. **Push your branch** to your forked repository.
5. **Create a pull request**.


## Contact

If you have any questions or feedback, feel free to reach out:

- [0x0060](https://0x0060.dev)
- [Email](mailto:ren@0x0060.dev)


## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/0x0060/ActiBot/blob/main/LICENSE) file for more information.

