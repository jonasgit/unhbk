# unhbk
Restore backupfile from Hogia Hemekonomi

På 1990-talet fanns ett program Hogia Hemekonomi. Det kunde göra backup av databasen till en fil med ändelsen *.HBK .
Filen går inte att läsa in direkt i Hogia Hemekonomi eller mitt wHHEK. Men detta program konverterar HBK-filen till en MDB-fil som går att använda av Hogia Hemekonomi eller wHHEK.
Programmet bör fungera både i Windows, MacOS och Linux men oftast vill man nog använda det i Windows eftersom MDB-filen inte går att använda någon annanstans.
Programmet körs från kommandoprompt/powershell och startas med att ange HBK-filen som argument. Exempel:
```
.\unhbk.exe 230731.HBK
```
