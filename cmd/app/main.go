// Arquivo principal do programa (entrypoint) 🫡
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

	"github.com/seu-usuario/meu-projeto-go/internal/hello"
	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
)

// Função principal do programa
func main() {
	// Mensagem inicial da aplicação
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")

	// Chamada para função de saudação
	hello.SayHello()
	
	n := 10
	seq := fibonacci.Fibonacci(n)
	ultimo := seq[len(seq)-1]

	fmt.Printf("F(%d) = %d\n", n, ultimo)
	fmt.Printf("Sequência de Fibonacci até F(%d): %v\n", n, seq)
}

