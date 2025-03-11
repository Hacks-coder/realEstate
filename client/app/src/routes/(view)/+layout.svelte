<script lang="ts">
	import Footer from '../../components/static/Footer.svelte';
	import Header from '../../components/static/Header.svelte';
	import Hero from '../../components/static/Hero.svelte';


	let { children } = $props();
    const heroBg = "/main.jpg"
</script>
<div style={`background-color: rgba(22, 22, 22, 0.5); background-image: url(${heroBg});`} class="bg-cover bg-center p-5 pt-10 pb-10">
	<Header />
	<Hero />
</div>
{@render children()}
<Footer />
