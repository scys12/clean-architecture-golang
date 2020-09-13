
mongo -- "$MONGO_INITDB_DATABASE" <<EOF
    db.category.insertMany( [
        {
            name: 'RAM'
        },
        {
            name: 'VGA'
        },
        {
            name: 'Processor'
        },
        {
            name: 'Motherboard'
        },
        {
            name: 'Storage'
        }
    ] );
EOF