# ARCANUS
ARCANUS is a customized payload generator/handler for penetration testing only.(Use at your own risk !).

For Assistance : egebal1996@gmail.com


# WHY USE ARCANUS ?
  İn pen.test world Metasploit is the mainstream for this job, but ARCANUS has few advantages.
  
- ARCANUS generates a unique payload for windows and linux systems that can't be detect with any antivirus programs. (Don't give any samples to Virus Total or similar web sites to keep it that way ;D )

- It has extra modules for exploitation. Ordinary reverse shell payloads offers only remote access to command prompts but ARCANUS has few special commands like " £persistence, £download, £upload, £meterpreter..."

- It is silent and continuous. Metasploit payloads attempts to connect remote host just for ones but when you execute ARCANUS payloads they makes connection attemps every 5 second silently in background.

- It is flexible. If you want to use it with Metasploit it has a meterpreter module for executeing meterpreter shellcodes on remote machine.

- Platform independent ! ARCANUS works both on windows and linux.


# HOW TO USE 

It works same as every reverse shell but it has some special module commands.
https://www.youtube.com/watch?v=BXYqeTs5RIE

   
                                                                                                     
                                                                                                     
                                                                                                     
                                                                                                     
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
                                                                                                     
                                                                                                     
                                                                                                  
# ANTIVIRU AWARENESS
  
  Please don't submit any payload samples to any antivirus sites or online forums. I will publish manual AV Scan detection scores continuously.
  


#/AV/    /Version/

- Last Scan : 13/03/2016
- Detection : 0/56


-- ALYac		20160313

-- AVG		20160313

-- AVware		20160313

-- Ad-Aware		20160313

-- AegisLab		20160313

-- Agnitum		20160313

-- AhnLab-V3		20160313

-- Alibaba		20160312

-- Antiy-AVL		20160313

-- Arcabit		20160313

-- Avast		20160313

-- Avira-(no-cloud)		20160313

-- Baidu		20160310

-- Baidu-International		20160313

-- BitDefender		20160313

-- Bkav		20160312

-- ByteHero		20160313

-- CAT-QuickHeal		20160312

-- CMC		20160307

-- ClamAV		20160311

-- Comodo		20160313

-- Cyren		20160313

-- DrWeb		20160313

-- ESET-NOD32		20160313

-- Emsisoft		20160313

-- F-Prot		20160313

-- F-Secure		20160313

-- Fortinet		20160313

-- GData		20160313

-- Ikarus		20160313

-- Jiangmin		20160313

-- K7AntiVirus		20160313

-- K7GW		20160313

-- Kaspersky		20160313

-- Malwarebytes		20160313

-- McAfee		20160313

-- McAfee-GW-Edition		20160313

-- eScan		20160313

-- Microsoft		20160313

-- NANO-Antivirus		20160313

-- Panda		20160313

-- Qihoo-360		20160313

-- Rising		20160313

-- SUPERAntiSpyware		20160313

-- Sophos		20160313

-- Symantec		20160310

-- Tencent		20160313

-- TheHacker		20160313

-- TrendMicro		20160313

-- TrendMicro-HouseCall		20160313

-- VBA32		20160313

-- VIPRE		20160313

-- ViRobot		20160313

-- Zillya		20160313

-- Zoner		20160313


-- nProtect		20160311
