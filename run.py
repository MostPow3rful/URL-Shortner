import os
import subprocess
import mysql.connector

from time import sleep
from sys import exit as error
from colorama import (
    Fore,
    init
)


class Color:
    RED = Fore.RED
    GREEN = Fore.GREEN
    CYAN = Fore.CYAN
    MAGENTA = Fore.MAGENTA
    WHITE = Fore.WHITE


class Config(Color):
    def __init__(self, _password: str = "", _distro: int = 0) -> None:
        self._clear()
        print(self._banner())
        sleep(1)
        self._USERNAME: str = "root"

        # Checking Password
        if _password is str():
            error(f"{Config.RED}[!] {Config.CYAN}Invalid Password For MySQL")
        self._PASSWORD: str = _password

        # Checking Distro
        if _distro not in range(1, 5):
            error(f"{Config.RED}[!] {Config.CYAN}Invalid Distro")
        self._DISTRO: str = _distro

        # Checking MySQL Package
        self._mysql_help() if self._check_mysql() is False else print(
            f"{Config.GREEN}[+] {Config.CYAN}MySQL Package Is Valid On Your System"
        )
        sleep(1)

        # Checking GoLang package
        self._golang_help if self._check_golang() is False else print(
            f"{Config.GREEN}[+] {Config.CYAN}GoLang Package Is Valid On Your System"
        )
        sleep(1)

        # Checking Directories And Files
        self._check_directory_file()

        # Configing MySQL
        self._config_mysql()

        # Creating Executable File Of main.go
        print(
            f"{Config.GREEN}[+] {Config.CYAN}Creating Executable File Of main.go"
        )
        os.system("go build ./main.go")

    def _banner(self) -> str:
        return f"""        
{Config.MAGENTA}_   _ ____  _       ____  _                _                   
{Config.GREEN}| | | |  _ \| |     / ___|| |__   ___  _ __| |_ _ __   ___ _ __ 
{Config.RED}| | | | |_) | |     \___ \| '_ \ / _ \| '__| __| '_ \ / _ \ '__|
{Config.CYAN}| |_| |  _ <| |___   ___) | | | | (_) | |  | |_| | | |  __/ |   
{Config.WHITE}\___/ |_| \_\_____| |____/|_| |_|\___/|_|   \__|_| |_|\___|_|                                                                  
        """

    def _check_mysql(self) -> bool:
        RESULT: str = subprocess.run(
            "command -v mysql",
            shell=True,
            capture_output=True,
        ).stdout.decode()

        return False if RESULT is str() else True

    def _check_golang(self) -> bool:
        RESULT: str = subprocess.run(
            "command -v go",
            shell=True,
            capture_output=True,
        ).stdout.decode()

        return False if RESULT is str() else True

    def _golang_help(self) -> None:
        match self._DISTRO:
            case 1:  # Debian
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With -> [dpkg , apt , ...]"
                )

            case 2:  # Arch
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With -> [pacman , yay , ...]"
                )

            case 3:  # Fedora
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With -> [dnf , yum , ...]"
                )

            case 4:  # Another
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install GoLang Package On Your System With Your Package Manager"
                )

    def _mysql_help(self) -> None:
        match self._DISTRO:
            case 1:  # Debian
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With -> [dpkg , apt , ...]"
                )

            case 2:  # Arch
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With -> [pacman , yay , ...]"
                )

            case 3:  # Fedora
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With -> [dnf , yum , ...]"
                )

            case 4:  # Another
                error(
                    f"{Config.RED}[!] {Config.CYAN}Please Install MySQL Package On Your System With Your Package Manager"
                )

    def _check_directory_file(self) -> None:
        DIRS: tuple = (
            "json",
            "log",
            "src",
            "src/config",
            "src/route",
            "src/structure",
            "static",
            "static/html",
            "static/css",
            "static/img",
            "static/js",
        )

        FILES: tuple = (
            "go.mod",
            "go.sum",
            "main.go",
            "log/log.log",
            "json/Secret.json",
            "src/config/config.go",
            "src/route/go.go",
            "src/route/register.go",
            "src/route/result.go",
            "src/route/root.go",
            "src/structure/data.go",
            "static/css/index.css",
            "static/css/result.css",
            "static/img/index_landing.png",
            "static/img/result_landing.png",
            "static/img/github.svg",
            "static/html/error.html",
            "static/html/index.html",
            "static/html/result.html",
            "static/js/index.js",
        )

        for directory in DIRS:
            error(
                f"{Config.RED}[-] {Config.CYAN}Invalid Directory ({directory})"
            ) if os.path.exists(directory) is False else print(
                f"{Config.GREEN}[+] {Config.CYAN}Valid Directory ({directory})"
            )
            sleep(1)

        for file in FILES:
            error(
                f"{Config.RED}[-] {Config.CYAN}Invalid File ({file})"
            ) if os.path.exists(file) is False else print(
                f"{Config.GREEN}[+] {Config.CYAN}Valid File ({file})"
            )
            sleep(1)

    def _config_mysql(self) -> None:
        database = mysql.connector.connect(
            host="localhost",
            user=self._USERNAME,
            password=self._PASSWORD
        )
        cursor = database.cursor()

        cursor.execute(
            f"CREATE database IF NOT EXISTS URLShortner"
        )
        database.commit()
        [... for _ in cursor]

        cursor.close()

        database = mysql.connector.connect(
            host="localhost",
            user="root",
            password=self._PASSWORD,
            database="URLShortner"
        )
        cursor = database.cursor()

        cursor.execute(
            "CREATE TABLE if not exists MYSQL_USER_PASS (username VARCHAR(300), password VARCHAR(300))"
        )

        cursor.execute(
            "INSERT INTO MYSQL_USER_PASS (username,password) VALUES (%s, %s)",
            (
                'root',
                self._PASSWORD
            )
        )
        database.commit()
        [... for _ in cursor]

        cursor.execute(
            "CREATE TABLE IF NOT EXISTS data (Title VARCHAR(300), URL VARCHAR(300), ID VARCHAR(300))"
        )
        database.commit()
        [... for _ in cursor]

        cursor.execute(
            "SELECT username, password FROM MYSQL_USER_PASS GROUP BY username, password"
        )
        RESULT = cursor.fetchall()[0]

        os.system('echo " " > json/Secret.json')
        with open(file="./json/Secret.json", mode="a") as file:
            file.write(
                '{\n\t"username":"'+RESULT[0] +
                '",\n\t"password":"'+RESULT[1]+'"\n}\n'
            )

    def _clear(self) -> None:
        os.system("clear")


def main() -> None:
    PASSWORD: str = input(
        f"""
        {Color.RED}[?] {Color.CYAN} Please Enter Your MySQL Password
        {Color.RED}[{Color.GREEN}Default=System Password{Color.RED}] {Color.WHITE}- {Color.RED}[{Color.GREEN}Need For Config{Color.RED}]
                    {Color.CYAN}-> """
    )

    try:
        KERNEL: int = int(
            input(f"""
        {Color.RED}[1] {Color.CYAN} Debian
        {Color.RED}[2] {Color.CYAN} Arch
        {Color.RED}[3] {Color.CYAN} Fedora
        {Color.RED}[4] {Color.CYAN} Another . . .
            
                    -> """)
        )
    except ValueError:
        error(
            f"{Color.RED}[!] {Color.CYAN}ValueError : U Must Enter Number In Range 1-4 In Kernel Option"
        )

    Config(
        PASSWORD,
        KERNEL
    )


if __name__ == "__main__":
    init(autoreset=True)
    main()
