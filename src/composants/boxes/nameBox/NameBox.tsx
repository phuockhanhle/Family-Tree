type boxInfo = {
    name: string
}

export const NameBox = (props: boxInfo) => {
    return (<span>{props.name}</span>);
}