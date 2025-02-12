<script>
	import Button from "../resuable/Button.svelte";


	import Text from "../resuable/Text.svelte";
    const aboutRecords = [
        {title: "50+", text: "Annual comback"},
        {title: "50+", text: "Annual comback"},
        {title: "50+", text: "Annual comback"},
        {title: "50+", text: "Annual comback"}
    ]
    const houseImage = [
        {img: "/house1.jpg", alt: "", size: "h-32"},
        {img: "house2.jpg", alt: "", size: "h-32"},
        {img: "/house3.jpg", alt: "", size: "h-54"}
    ]

</script>
<section class="mt-20 flex flex-col items-center">
    <div class="text-center">
        <Text text = "ABOUT US" className = "font-bold text-xl" />
        <Text text = "We offer the best price you can thing of" />
        <Button text = "Check out" className = "bg-standard rounded-full w-28 p-3 mt-5" />
    </div>
    <div class="grid grid-cols-2 grid-rows-2 gap-10 mt-10">
        {#each aboutRecords as item}
            <div>
                <Text text = {item.title} className = "text-3xl font-bold" />
                <Text text = {item.text} className = "text-sm" />
            </div>
        {/each}
    </div>
    <div class="flex gap-2 mt-10">
        {#each houseImage as item}
            <div class={` w-28 bg-cover bg-center ${item.size}  rounded-lg md:w-48 md:`} style={`background-image: url(${item.img})`}>
            </div>
        {/each}
    </div>
</section>