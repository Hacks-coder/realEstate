<script>
	import Button from "../resuable/Button.svelte";


	import Text from "../resuable/Text.svelte";
    const heroBg = "/main.jpg"
</script>
<section  class="grid grid-cols-1 place-items-center md:grid-cols-2 mt-10 md:mt-20 lg:mt-32">
    <div class="flex flex-col items-center">
        <div class="text-center w-64 md:w-96 lg:w-fit">
            <Text text = "Let’s find your comfort house" className = "text-2xl md:text-3xl lg:text-6xl" />
            <Text text = "Let’s find the perfect place where comfort meets convenience, and you can truly feel at home" className = "text-xs" />
        </div>
        <div class="flex gap-5 mt-5 md:mt-16">
            <Button text = "Get Started" />
            <Button text = "Get Started" />
        </div>

    </div>
    <div style={`background-image: url(${heroBg});`} class="w-72 h-64 bg-cover bg-center     shadow-sm shadow-black mt-10">

    </div>
</section>