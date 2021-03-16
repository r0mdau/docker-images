package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func GetLocale() string {
  envlang, ok := os.LookupEnv("LANG")
  if ok {
    return strings.Split(envlang, ".")[0]
  }

  cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
  output, err := cmd.Output()
  if err == nil {
    return strings.Trim(string(output), "\r\n")
  }

  return "en"
}

func main() {
    locale := GetLocale()

    coverFR := `
Bonjour Lightspeed !

Super content que vous soyez arrivé jusque ici. J'adore les challenges (CTF sécurité par exemple...) mais on va s'arrêter là.

Je suis arrivé récemment à Montréal et suis disponible pour travailler avec vous dès le 1er Juin.

Une lettre de motivation qui est originale ce n'est pas évident, je dois avoir mis une petite touche, mais pas trop, prochain coup
je développe un OS !

Mon CV (https://cv.romaindauby.fr) est pas mal détaillé, vous y trouverez toutes les informations écrites d'une façon condensée me concernant. Tout le reste pour apprendre à mieux me connaître... j'ai hâte de vous le dire de vive voix.

Le code de cette lettre de motivation est disponible ici : https://github.com/r0mdau/docker-images/tree/master/coverletter/lightspeed
`

    coverEN := `
Hello Lightspeed!

Super glad you made it this far. I love challenges (CTF security for example ...) but we will stop there.

I recently arrived in Montreal and am available to work with you on June 1st.

A cover letter that is original is not easy, I must have put a little touch, but not too much, next move
I will develop an OS!

My CV (https://cv.romaindauby.fr) is quite detailed, you will find all the information written in a condensed way about me. Everything else to get to know me better ... I can't wait to tell you face to face.

The code for this cover letter is available here: https://github.com/r0mdau/docker-images/tree/master/coverletter/lightspeed
`

    if locale == "fr_FR" || locale == "fr_CA"{
        fmt.Println(coverFR)
    } else {
        fmt.Println(coverEN)
    }
}

