package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net"
    "net/http"

    "github.com/dgrijalva/jwt-go"
)


type MtProtoProxy struct {
    secretKey string
}


func NewMtProtoProxy(secretKey string) *MtProtoProxy {
    return &MtProtoProxy{secretKey: secretKey}
}


func (p *MtProtoProxy) Listen() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go p.handleConnection(conn)
    }
}


func (p *MtProtoProxy) handleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            log.Fatal(err)
        }

        msg := mtprotoMsg{}
        err = json.Unmarshal(buf[:n], &msg)
        if err != nil {
            log.Println(err)
            continue
        }

        switch msg.Cmd {
        case "auth_key":
            p.handleAuthKey(conn, msg)
        case "mtproto_msg":
            p.handleMtprotoMsg(conn, msg)
        default:
            log.Println("unknown command", msg.Cmd)
        }
    }
}


func (p *MtProtoProxy) handleAuthKey(conn net.Conn, msg mtprotoMsg) {
    authKey := msg.Data.(string)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
        IssuedAt: time.Now().Unix(),
    })

    tokenString, err := token.SignedString([]byte(p.secretKey))
    if err != nil {
        log.Fatal(err)
    }

    conn.Write([]byte(`{"result": "ok", "auth_key": "` + tokenString + `"}`))
}


func (p *MtProtoProxy) handleMtprotoMsg(conn net.Conn, msg mtprotoMsg) {

    var decodedMsg mtprotoMsg
    err := json.Unmarshal(msg.Data, &decodedMsg)
    if err != nil {
        log.Println(err)
        return
    }

    switch decodedMsg.Cmd {
    case "get_me":
        result := mtprotoGetMeResult{
            User: &tgUser{
                ID:       123456,
                FIRST_NAME: "John",
                LAST_NAME: "Doe",
            },
            COUNT: 1,
        }
        json.NewEncoder(conn).Encode(&result)
    case "get_dialogs":
        result := mtprotoGetDialogsResult{
            Diags: []*tgDialog{},
            COUNT: 1,
        }
    json.NewEncoder(conn).Encode(&result)
    default:
log.Println("unknown command", decodedMsg.Cmd)
}
}

type mtprotoMsg struct {
    Cmd   string      `json:"cmd"`
    Data  json.RawMessage `json:"data"`
    Retry int         `json:"retry"`
}

type mtprotoGetMeResult struct {
    User *tgUser    `json:"user"`
    COUNT int        `json:"count"`
}

type mtprotoGetDialogsResult struct {
    Diags []*tgDialog `json:"diags"`
    COUNT int        `json:"count"`
}

type tgUser struct {
    ID       int     `json:"id"`
    FIRST_NAME string `json:"first_name"`
    LAST_NAME string `json:"last_name"`
}

type tgDialog struct {
    ID   int     `json:"id"`
    TITLE string `json:"title"`
}
