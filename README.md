# ARCANUS
ARCANUS is a customized payload generator/handler for penetration testing only.(Use at your own risk !).



# WHY USE ARCANUS ?
  İn pen.test world Metasploit is the mainstream for this job, but ARCANUS has few advantages.
  
- ARCANUS generates a unique payload for windows and linux systems that can't be detect with any antivirus programs. (Don't give any samples to Virus Total or similar web sites to keep it that way ;D )

-It has extra modules for exploitation. Ordinary reverse shell payloads offers only remote access to command prompts but ARCANUS has few special commands like " £persistence, £download, £upload, £meterpreter..."

-It is silent and continuous. Metasploit payloads attempts to connect remote host just for ones but when you execute ARCANUS payloads they makes connection attemps every 5 second silently in background.

-It is flexible. If you want to use it with Metasploit it has a meterpreter module for executeing meterpreter shellcodes on remote machine.


# HOW TO USE 

It works same as every reverse shell but it has some special module commands.

   
                                                                                                     
                                                                                                     
                                                                                                     
                                                                                                     
     [ COMMAND ]                                       [DESCRIPTION]                                 
                            
                                                                                                     
     (*) £METERPRETER -C:                              This command executes given powershell        
                                                         meterpreter shellcode for metasploit        
                                                          integration.                               
                                                                                                     
                                                                                                     
     (*) £PERSISTENCE:                                 This command installs a persistence module    
                                                         to remote computer for continious acces.    
                                                                                                     
                                                                                                     
     (*) £METASHELL:                                   This command executes given metasploit reverse
                                                         shell shellcode on remote machine.          
                                                                                                     
                                                                                                     
     (*) £UPLOAD -F "filename.exe":                    This command uploads a choosen file to        
                                                         remote computer via tcp socket stream.      
                                                                                                     
                                                                                                     
     (*) £UPLOAD -G "http://filepath/filename.exe":    This command uploads a choosen file to        
                                                         remote computer via http get method.        
                                                                                                     
                                                                                                     
     (*) £DOWNLOAD -F "filename.exe":                  This command download a choosen file          
                                                         from remote computer via tcp socket stream. 
                                                                                                     
                                                                                                     
     (*) £REQUEST -A "www.site.com":                   This command sends a http get request on      
                                                         remote machine to given addres.             
                                                                                                     
                                                                                                     
     (*) £PLEASE "any command":                        This command asks users comfirmation for      
                                                         higher privilidge operations.               
                                                                                                     
                                                                                                     
     (*) £DESKTOP                                      This command adjusts remote desktop options   
                                                         for remote connection on target machine     
                                                                                                     
                                                                                                     
  
