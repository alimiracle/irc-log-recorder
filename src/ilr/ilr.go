/*
* irc log recorder
* irc bot to  Watching Everything that happens in the irc channel  and record it in log and its have web server To enable users to access the logs from Web Browser and telnet
* Copyright (c) 2016 ali abdul ghani <blade.vp2020@gmail.com>
*    This Program is free software: you can redistribute it and/or modify
*    it under the terms of the GNU  General Public License as published by
*    the Free Software Foundation, either version 3 of the License, or
*    (at your option) any later version.
*    This Program is distributed in the hope that it will be useful,
*    but WITHOUT ANY WARRANTY; without even the implied warranty of
*    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*    GNU General Public License for more details.
 *    You should have received a copy of the GNU General Public License
*    along with this library.  If not, see <http://www.gnu.org/licenses/>.
*/

package main
import (
"encoding/json"
    "github.com/thoj/go-ircevent"
    "fmt"
"time"
"os"
"net/http"
"io"
 "io/ioutil"
 "log"
)
type CONFIG struct {
Server string
Name string
User string
Password string
RoomName string
}

func copy(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
 log.Fatal(err)
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
    return
}

func handler(w http.ResponseWriter, r *http.Request) {
copy("/var/log/irclog.txt", "temp.txt")
file, err := os.Open("temp.txt")
 if err != nil {
 log.Fatal(err)
 }
 data, err := ioutil.ReadAll(file)
 if err != nil {
 log.Fatal(err)
 }
s := string(data)
    fmt.Fprintf(w,s)
file.Close()
os.Remove("temp.txt")
}

func server() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func main() {
go server()
file, err := os.Open("/etc/ircconfig.txt")
 if err != nil {
 log.Fatal(err)
 }
var all CONFIG

 data, err := ioutil.ReadAll(file)
 if err != nil {
 log.Fatal(err)
 }
rede := json.Unmarshal(data, &all)
    if rede != nil {
 log.Fatal(rede)
}
f, err := os.Create("/var/log/irclog.txt")
if err != nil {
 log.Fatal(err)
}
    con := irc.IRC(all.Name, all.User)
con.Password = all.Password

    er := con.Connect(all.Server)
    if er != nil {
 log.Fatal(er)
}
    con.AddCallback("001", func (e *irc.Event) {
        con.Join(all.RoomName)    
    })
    con.AddCallback("JOIN", func (e *irc.Event) {
f.WriteString(e.Nick)
f.WriteString(" is joined in ")
f.WriteString(all.RoomName)
f.WriteString("\n")
    })
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
f.WriteString(e.Nick)
t := time.Now()
f.WriteString(" ")
f.WriteString(t.String())
f.WriteString(" ")
f.WriteString(e.Message())
f.WriteString("\n")
    })
      con.Loop()
}
