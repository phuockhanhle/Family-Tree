import { Box } from "@mui/material";
import axios from "axios";
import { useEffect, useState } from "react";
import Tree from "react-d3-tree/lib/Tree";
import { RawNodeDatum, TreeNodeDatum } from "react-d3-tree/lib/types/common";
import { AddFamilyModal } from "../../components/addFamily/AddFamilyModal";

const width = '100vw';
const height = 300;
const circleX = '50vw';
const circleY = height / 2;
const circleRadius = 50;
const initialMousePosition = {
    x: circleX,
    y: height / 2
}
const url = 'https://gist.githubusercontent.com/phuockhanhle/73d2926eaf7e9f331ec55ce66c4ffdf6/raw/example_tree.json'

export const DashBoard = () => {
    const [tree, setTree] = useState<RawNodeDatum | RawNodeDatum[]>({
        name: "Root",
        attributes: {id: '0'},
        children: []
    });
    const [node, setNode] = useState<undefined | TreeNodeDatum>(undefined);
    const [data, setData] = useState<RawNodeDatum | RawNodeDatum[]>([])

    useEffect(() => {
        axios.get(url).then(response => {
            console.log(response.data);
            const convertedDataTree = parseDataToRawNodeDatum(response.data);
            setData(convertedDataTree);
        });
    }, [])

    const close = () => setNode(undefined);
    const handleSubmitName = (name: string, id: string) => {
        if (node) {
            const newTree = bfs(node.name, tree, name, id);
            if (newTree) {
                setTree(newTree);
            }
        }
    };
    const handleClickAddFamily = (familyTree: RawNodeDatum) => {
        if (node) {
            let newTree = bfs_tree(node.name, tree, familyTree);
            console.log(newTree)
            if (newTree) {
                setTree(newTree);
            }
        }
    }

    return (
        <>
            <Box sx={{
                width: '100vw',
                height: '900px'
            }}>
                <Tree data={tree} translate={{ x: 200, y: 300 }} onNodeClick={(datum) => { setNode(datum.data) }} />
                <AddFamilyModal
                    isOpen={Boolean(node)}
                    onClose={close}
                    onSubmit={handleSubmitName}
                    onClickAddFamily={handleClickAddFamily}
                    familyTree={data} />

            </Box>

        </>


    );
}

function bfs(name: string, tree: RawNodeDatum | RawNodeDatum[], newNodeName: string, newNodeId: string) {
    const queue: RawNodeDatum[] = [];
    queue.unshift(tree as RawNodeDatum);
    while (queue.length > 0) {
        const curNode = queue.pop();
        if (curNode && curNode.children) {
            if (curNode.name === name) {
                curNode.children.push({
                    name: newNodeName,
                    attributes: {id: newNodeId},
                    children: []
                });
                return { ...tree };
            }
            const len = curNode.children.length;

            for (let i = 0; i < len; i++) {
                queue.unshift(curNode.children[i]);
            }
        }



    }
}

function bfs_tree(name: string, tree: RawNodeDatum | RawNodeDatum[], newTree: RawNodeDatum) {
    const queue: RawNodeDatum[] = [];
    queue.unshift(tree as RawNodeDatum);
    while (queue.length > 0) {
        const curNode = queue.pop();
        if (curNode && curNode.children) {
            if (curNode.name === name) {
                curNode.children.push(newTree);
                return { ...tree };
            }
            const len = curNode.children.length;

            for (let i = 0; i < len; i++) {
                queue.unshift(curNode.children[i]);
            }
        }
    }

}

function parseDataToRawNodeDatum(data: any) {
    // init tree
    let tree: RawNodeDatum = {
        name: data[0].end.properties.FirstName,
        attributes: {id: String(data[0].end.identity)},
        children: []
    }
    
    //parsing tree
    for (let i = 0; i < data.length; i++) {
        const queue: RawNodeDatum[] = [];
        queue.unshift(tree as RawNodeDatum);
        while (queue.length > 0) {
            const curNode = queue.pop();
            if (curNode && curNode.children && curNode.attributes) {
                if (String(data[i].end.identity) === curNode.attributes.id) {
                    curNode.children.push({
                        name: data[i].start.properties.FirstName,
                        attributes: {id: String(data[i].start.identity)},
                        children: []
                    });
                }
                const len = curNode.children.length;
                for (let j = 0; j < len; j++) {
                    queue.unshift(curNode.children[j]);
                }
            }
        }
    }
    console.log(tree);
    return tree;

}