<script>

	import Text from "../resuable/Text.svelte";
    const testimonies = [
        {
            img: "/tes1.jpg",
            title: "Daniel I",
            testimony: "We offer the best price you can thingk of"
        },
        {
            img: "/tes2.jpg",
            title: "Daniel I",
            testimony: "We offer the best price you can thingk of"
        },
    ]

</script>
<section class="mt-20">
    <Text text = "Testimonies" className = "text-2xl font-bold" />
    <Text text = "See what people say about us" className = "" />
    <div class="grid gap-2 grid-cols-[170px_170px] place-items-center mt-5 md:grid-cols-[200px_200px]">
        {#each testimonies as item}
            <div class="text-center flex flex-col py-5 items-center shadow-md shadow-black p-2 rounded-lg mt-5">
                <div class="w-10 h-10 rounded-full border-2 border-white overflow-hidden">
                    <img src={item.img} alt="">
                </div>
                <Text text = {item.title} className = "font-bold text-xl" />
                <Text text = {item.testimony} className = "text-xxs" />
            </div>
        {/each}
    </div>
</section>