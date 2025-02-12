<script>
    const companiesOne = [
        {logo: "/google.png", alt: ""},
        {logo: "/amazon.png", alt: ""},
        {logo: "spotify.png", alt: ""},
        {logo: "/apple.png", alt: ""},
        {logo: "/microsoft.png", alt: ""}
    ]
    const compainesTwo = [
        {logo: "/pepsi.png", alt: ""},
        {logo: "/adidas.png", alt: ""}
    ]
</script>


<section class="flex flex-col items-center mt-10">
    <div class="flex">
        {#each companiesOne as item}
        <img src={item.logo} alt={item.alt} class="w-16 md:w-32">
        {/each}
    </div>
    <div class="flex">
        {#each compainesTwo as item}
            <img src={item.logo} alt={item.alt} class="w-16 md:w-32">
        {/each}
    </div>
</section>