# UnityProjectCleaner - Delete Library directory in the Unity project

```
> UnityProjectCleaner.exe -h  
  usage: 
      UnityProjectCleaner.exe [-d] targetDirectory 

  options:
      -d  dry run mode
      -y  automatic yes to prompt

  examples:
      UnityProjectCleaner.exe /home/test/unity_projects/

> UnityProjectCleaner.exe c:\work\    

Do you want to start the clean up process? (y/n) [n]: y
Start clean up process...unity project found...path=c:\work\UnityProjectA
   delete... c:\work\UnityProjectA\Library
   delete... c:\work\UnityProjectA\Logs
   delete... c:\work\UnityProjectA\obj
   delete... c:\work\UnityProjectA\Temp
unity project found...path=c:\work\UnityProjectB
   delete... c:\work\UnityProjectB\Library
   delete... c:\work\UnityProjectB\Logs
   delete... c:\work\UnityProjectB\obj
   delete... c:\work\UnityProjectB\Temp
unity project found...path=c:\work\UnityProjectC
   delete... c:\work\UnityProjectC\Library
   delete... c:\work\UnityProjectC\Logs
   delete... c:\work\UnityProjectC\obj
   delete... c:\work\UnityProjectC\Temp
```


## Copyright and license
Copyright (c) 2022 yoggy

Released under the [MIT license](LICENSE)