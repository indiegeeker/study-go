package main 

type tree_data struct{
	id int 
	title string
	children []*tree_data
}

func main(){
	/*
	let tree = [
    {
        id: 1,
        title: "grandfather",
        children: [
            {
                id: 3,
                title: "father"
            },
            {
                id: 4,
                title: "father2",
                children: [{
                    id: 6,
                    title: "child1"
                }]
            }
        ]
    },
    {
        id: 2,
        name: "grandfather2",
        children: [{
            id: 5,
            name: "father3"
        }]
    }
];


let list = [
    { id: 1, title: "grandfather", path: ['grandfather'] },
    { id: 3, title: "father", path: ['grandfather', 'father'] },
    { id: 4, title: "father2", path: ['grandfather', 'father2'] },
    { id: 6, title: "child1", path: ['grandfather', 'father2', 'child1'] },
    { id: 2, title: "grandfather", path: ['grandfather2'] },
    { id: 5, title: "father3", path: ['grandfather2', 'father3'] },
];
	*/
	data = make([]*tree_data,0)
	data := [
		{
			id: 1,
			title: "grandfather",
			children: [
				{
					id: 3,
					title: "father"
				},
				{
					id: 4,
					title: "father2",
					children: [{
						id: 6,
						title: "child1"
					}]
				}
			]
		},
		{
			id: 2,
			name: "grandfather2",
			children: [{
				id: 5,
				name: "father3"
			}]
		}
	]
	
	fmt.Println(data)
}