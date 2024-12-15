Here’s the updated `README.md` with the additional sections and emojis to make it more engaging:

---

# LoopRun 🚀

LoopRun is a simple CLI tool that allows you to run and restart your commands easily. It listens for keyboard input and allows you to restart your command (e.g., a server) by pressing `Ctrl+R`. It's perfect for local development environments where you need to restart servers frequently.

## Features ✨

- Run any command and restart it with `Ctrl+R` 🔄
- Automatically forwards the command's output to your terminal 📜
- Handles stopping the command gracefully on exit 🚪

## Installation 💻

To install LoopRun using Go, run the following command:

```bash
go install github.com/katistix/looprun@latest
```

Make sure your `$GOPATH/bin` is part of your system's `PATH` so that the `looprun` command is available globally.

## Usage 🎮

You can use LoopRun to run any command. For example, you can use it to start a simple Python HTTP server that you can restart with a key press.

### Example:

Start a Python HTTP server that serves files from the current directory:

```bash
looprun python3 -m http.server
```

Once started, you can press `Ctrl+R` to restart the server, or `Ctrl+C` to quit the program.

### Usage Syntax:

```bash
looprun <command> [args...]
```

## Contributing 🤝

Contributions are welcome! If you'd like to improve this project, feel free to open an issue or a pull request. Here's how you can get started:

1. Fork this repository 🍴
2. Create a new branch (`git checkout -b feature-branch`) 🌿
3. Make your changes and commit them (`git commit -m 'Add a new feature'`) 💡
4. Push to the branch (`git push origin feature-branch`) 🚀
5. Open a pull request 📩

Let's build something awesome together! 🙌

## License 📜

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Support ☕️

If you like this project and want to support its development, you can buy me a coffee:

<a href="https://www.buymeacoffee.com/katistix" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/yellow_img.png" alt="Buy Me A Coffee"></a>


Thank you for your support! 😊