<script>
	
	import { allTodos} from './rpc.svelte'
	
	
	let todos = allTodos()

</script>

<main>

	<div style="background-color: white;color:black;">
		{#await todos}
		<p>...waiting</p>
	{:then data}
		<p>Todos {JSON.stringify(data)}</p>
	{:catch error}
		<p style="color: red">{error.message}</p>
	{/await}

	</div>


	
</main>

<style>

	main {
		height: 100%;
		width: 100%;
	}

	#result {
		margin-top: 1rem;
		font-size: 1.5rem;
	}

	button {
		-webkit-appearance: default-button;
		padding: 6px;
	}

	#name {
		border-radius: 3px;
		outline: none;
		-webkit-font-smoothing: antialiased;
	}

	#logo {
		width: 40%;
		height: 40%;
		padding-top: 20%;
		margin: auto;
		display: block;
		background-position: center;
		background-repeat: no-repeat;
		background-image: url("assets/images/logo-dark.svg");
	}

</style>