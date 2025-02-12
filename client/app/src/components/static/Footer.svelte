<script>
	import Button from "../resuable/Button.svelte";


	import Text from "../resuable/Text.svelte";
    const performanceScores = [
        {title: "50+", text: "Annual comback"},
        {title: "50+", text: "Annual comback"},
        {title: "50+", text: "Annual comback"}
    ]
</script>
<footer class="bg-standard mt-20 p-5 grid grid-cols-1 md:grid-cols-2">
    <div class="flex flex-col gap-5">
        <div>
            <Text text = "We’re revolutuonizing the landscape" className = "font-bold text-xl md:text-2xl" />
        </div>
        <div>
            <Text text = "We have a strong commitment to collaborate and effectively communicate. Your can get to us with just a click of this button. let’s talk" className = "text-xs md:text-xl" />
            <Button text = "Talk to us"  />
        </div>
    </div>
    <div class="mt-10 flex gap-3 justify-center">
        {#each performanceScores as item}
            <div>
                <Text text = {item.title} className = "text-2xl font-bold" />
                <Text text = {item.text} className = "text-sm" />
            </div>
        {/each}
    </div>
</footer>