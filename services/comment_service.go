package services

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

func CollectComments(filePath string) (map[string]string, error) {

	resultMap := make(map[string]string)
	regex := regexp.MustCompile(`@(\w+)\s+(.+)`)

	// Créer un ensemble de fichiers token
	fs := token.NewFileSet()

	// Parser le fichier pour construire un arbre syntaxique (AST)
	node, err := parser.ParseFile(fs, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Parcourir l'AST pour trouver les fonctions et leurs commentaires
	ast.Inspect(node, func(n ast.Node) bool {
		// Vérifier si le nœud est une déclaration de fonction
		funcDecl, ok := n.(*ast.FuncDecl)
		if ok {
			// Récupérer le nom de la fonction
			funcName := funcDecl.Name.Name

			if funcName == "main" && funcDecl.Doc != nil {
				// Parcourir les commentaires de la fonction
				for _, comment := range funcDecl.Doc.List {
					// Vérifier si le commentaire correspond au format attendu
					if regex.MatchString(comment.Text) {
						// Extraire les informations du commentaire
						matches := regex.FindStringSubmatch(comment.Text)
						// Stocker les informations dans la map
						resultMap[strings.TrimSpace(matches[1])] = strings.TrimSpace(matches[2])
					}
				}
			}
		}
		return true
	})

	return resultMap, nil
}
