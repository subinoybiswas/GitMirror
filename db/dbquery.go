package db

import (
	"database/sql"
	"fmt"
)

func getUserMappingByUsername(db *sql.DB, username string) ([]UserMapping, error) {
    query := "SELECT id, github_username, gitlab_username FROM username_mappings WHERE github_username = ? OR gitlab_username = ?"
    rows, err := db.Query(query, username, username)
    if err != nil {
        return nil, fmt.Errorf("failed to execute query: %v", err)
    }
    defer rows.Close()

    var mappings []UserMapping

    for rows.Next() {
        var mapping UserMapping

        if err := rows.Scan(&mapping.ID, &mapping.GitHubUsername, &mapping.GitLabUsername); err != nil {
            return nil, fmt.Errorf("error scanning row: %v", err)
        }

        mappings = append(mappings, mapping)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error during rows iteration: %v", err)
    }

    return mappings, nil
}

func FindService(db *sql.DB, username string) ([]UserMapping, error) {
    userMappings, err := getUserMappingByUsername(db, username)
    if err != nil {
        return  userMappings,err
    }
    return userMappings,nil
}
