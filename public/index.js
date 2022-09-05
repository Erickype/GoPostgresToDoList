function updateDb(item) {
    let input = document.getElementById(item)
    let newItem = input.value
    fetch(`/update?oldItem=${item}&newItem=${newItem}`, {method: "PUT"}).then(res =>{
        if (res.status === 200){
            alert("Database updated")
            //window.location.pathname = "/"
        }
    })
}