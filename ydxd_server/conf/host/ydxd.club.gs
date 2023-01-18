server: {
	listen: 4000
	host: localhost

	location: {
		url: /
		type: static
		http_dir: `/Users/wxj/workplace/5.Work/5.yidianxinde/ydxd/ydxd_docs`
		# prefix: /static
	}
}

server: {
	listen: 4001
	host: localhost

	location: {
		url: /
		type: static
		http_dir: `/Users/wxj/workplace/5.Work/5.yidianxinde/ydxd/ydxd_manager/build`
		# prefix: /static
	}
}

