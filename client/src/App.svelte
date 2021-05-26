<script>
	let setup = "Setup"
	let punchline = "Punchline"

	async function fetch_post() {
		const proxy = "CORS_PROXY_URL"
		const res = await fetch(proxy + 'https://tavuli.com/api/init', {
			method: 'POST',
			headers:{
				'Content-Type':'application/json'
			},
			body: JSON.stringify({
				token:'qC5Ym%2BBvXzabGp34jJlKvnC6wCrr3pLCwBzsLoSzl4k%3D'
			})
		})
		const json = await  res.json()
		if(json.status === 202){
			setup = json.token
			punchline = json.company
		}
	}
	async function fetch_get(){
		const proxy = "CORS_PROXY_URL"
		const url = "https://joke-api-strict-cors.appspot.com/random_joke"
		const res = await fetch(proxy + url)
		const json = await  res.json()
		setup = json.setup
		punchline = json.punchline
	}
</script>

<main>
	<h1>Cors Demo</h1>
	<div>{setup}</div>
	<button on:click={fetch_post}>Fetch POST</button>
	<button on:click={fetch_get}>Fetch GET</button>
	<div>{punchline}</div>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>