# Website Monitor

This is a website monitoring tool written in Go that checks whether a specific website is online or not.


## Installation

After cloning the repository to your device, navigate to the `/src` folder in your Go directory and place the `monitora-site` folder inside it.

```
<your-local-disk>/Go/src/monitora-site/
```

## Execution

To run the project, open the terminal in the directory where it is located and execute the following command:

```shell
go run monitora-site.go
```

## Additional Information

To monitor websites, you need to write the URLs of the websites you want to monitor in the `sites.txt` file using the following format:

```
http(s)://www.example.com/org/etc
```

If you want to monitor multiple websites, add a new line and include the URL of the next website. For example:

```
https://www.google.com
https://www.freecodecamp.org/learn/
```

## Contributing

Contributions are welcome! If you find any bugs or want to add new features, please feel free to submit an issue or create a pull request.
