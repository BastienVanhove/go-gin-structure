package main

import (
    contextManage "root/Context/contextManage"
   "root/Core/envRead"
)

func main() {
    mapEnv := envRead.Read()
    contextManage.Start(mapEnv)
}
