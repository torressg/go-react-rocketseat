package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Obter o diretório atual de trabalho
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Construir o caminho absoluto para o arquivo de configuração
	configPath := filepath.Join(workingDir, "internal/store/pgstore/migrations/tern.conf")

	// Verificar se o arquivo de configuração existe
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Arquivo de configuração não encontrado:", configPath)
		return
	}

	// Configurar o comando
	cmd := exec.Command("tern", "migrate", "--migrations", filepath.Join(workingDir, "internal/store/pgstore/migrations"), "--config", configPath)

	// Capturar a saída padrão e de erro
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Executar o comando
	if err := cmd.Run(); err != nil {
		fmt.Println("Erro ao executar o comando:", err)
	}
}
