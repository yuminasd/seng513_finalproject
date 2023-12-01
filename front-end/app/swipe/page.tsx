

export default function Swipe() {
    
    const groups = ['group 1', 'group 2', 'test'];
    
    
   
    
    return (
        <main>
            <select name="selectgroupbox" >
                <option value="" selected disabled hidden>Group Name</option>
                {groups.map((e, key) => {
                    return <option key={key} value={e}>{e}</option>;
                })}
            </select>
        </main>
    )
}
