<script>
	import Router from 'svelte-spa-router';
	import {wrap} from 'svelte-spa-router/wrap'
	import { Row, Container, Col } from 'sveltestrap'
	import Navigation from './componets/Navigation.svelte'
	import Footer from './componets/Footer.svelte'
	import Home from './pages/home/Home.svelte'
	import Company from './pages/company/company.svelte'
	import Login from './pages/Login.svelte'
	import NotFound from './pages/Notfound.svelte'

	let token = "123123"
	let routes = ""
	let isNav = false
	if(token === ""){
		routes = {
			'/': wrap({
				component:Login
			}),
			'*': NotFound,
		}
	}else{
		isNav = true
		routes = {
			'/': wrap({
				component:Home
			}),
			'/company': wrap({
				component:Company
			}),
			'*': NotFound,
		}
	}

	
</script>
{#if isNav}
<Navigation />
{/if}
<Container style="margin-top: 70px;">
	<Row>
		<Col>
			<Router routes={routes} />
			
		</Col>
	</Row>
	
</Container>
<Footer />