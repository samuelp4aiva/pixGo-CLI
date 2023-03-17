package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func addCmd() *cobra.Command {
	var key, alias string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Adiciona uma chave PIX",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := setupDatabase()
			if err != nil {
				return err
			}

			err = addPixKey(db, key, alias)
			if err != nil {
				return err
			}

			fmt.Printf("Chave PIX adicionada com sucesso: %s (alias: %s)\n", key, alias)
			return nil
		},
	}

	cmd.Flags().StringVarP(&key, "key", "k", "", "Chave PIX (obrigatório)")
	cmd.Flags().StringVarP(&alias, "alias", "a", "", "Apelido para a chave PIX (opcional)")
	cmd.MarkFlagRequired("key")
	return cmd
}

func listCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lista todas as chaves Pix",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := setupDatabase()
			if err != nil {
				return err
			}

			pixKeys, err := listPixKeys(db)
			if err != nil {
				return err
			}

			fmt.Println("Chaves Pix: ")
			for _, pixKey := range pixKeys {
				fmt.Printf("- %s (alias: %s)\n", pixKey.Key, pixKey.Alias)
			}

			return nil
		},
	}

	return cmd
}