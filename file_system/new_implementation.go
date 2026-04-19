Node {
    id string (PK)

    name string
    type ENUM (FILE, DIRECTORY)

    parent_id string   // tree structure
    owner_id string

    size long          // file size or cached dir size

    is_deleted boolean

    created_at timestamp
    updated_at timestamp
}
INDEX(parent_id)
INDEX(owner_id)


FileContent {
    node_id string (PK, FK)

    storage_url string   // S3 / blob storage
    size long
}

AccessControl {
    id string (PK)

    node_id string
    user_id string

    permission ENUM (READ, WRITE, OWNER)
}
INDEX(user_id, node_id)


SearchIndex {
    node_id
    name
    owner_id
}

