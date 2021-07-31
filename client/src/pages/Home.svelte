<script>
    import Panel from '../componets/Panel.svelte'

    const queryString = window.location.search
	const urlParams = new URLSearchParams(queryString)
	const token_browser = urlParams.get('token')
	let client_device = "WEBSITE"
</script>

<Panel />