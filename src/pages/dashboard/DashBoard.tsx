import { Box, Button } from "@mui/material";
import axios from "axios";
import { useEffect, useState } from "react";
import Tree from "react-d3-tree/lib/Tree";
import { RawNodeDatum, TreeNodeDatum } from "react-d3-tree/lib/types/common";
import { AddFamilyModal } from "../../components/addFamily/AddFamilyModal";
import { v4 as uuidv4 } from 'uuid';
import { useNavigate } from "react-router-dom";
import { logOut } from '../../components/utils/HelperFuncs';
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
    const navigate = useNavigate();
    const [tree, setTree] = useState<RawNodeDatum | RawNodeDatum[]>({
        name: "Root",
        attributes: {identity: 0},
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
    const handleSubmitName = (name: string, identity: number) => {
        if (node && node.attributes) {
            const newTree = bfs(Number(node.attributes.identity), tree, name, identity);
            if (newTree) {
                setTree(newTree);
            }
        }
    };
    const handleClickAddFamily = (familyTree: RawNodeDatum) => {
        if (node) {
            console.log(familyTree);
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
                height: '500px'
            }}>
                <Tree data={tree} translate={{ x: 200, y: 300 }} onNodeClick={(datum) => { setNode(datum.data) }} />
                <AddFamilyModal
                    isOpen={Boolean(node)}
                    onClose={close}
                    onSubmit={handleSubmitName}
                    onClickAddFamily={handleClickAddFamily}
                    familyTree={data} />
            </Box>
            <Button variant="contained" onClick={() => logOut(navigate)}> Logout </Button>

        </>


    );
}

function bfs(identity: number, tree: RawNodeDatum | RawNodeDatum[], newNodeName: string, newNodeId: number) {
    const queue: RawNodeDatum[] = [];
    queue.unshift(tree as RawNodeDatum);
    while (queue.length > 0) {
        const curNode = queue.pop();
        if (curNode && curNode.children && curNode.attributes) {
            if (curNode.attributes.identity === identity) {
                curNode.children.push({
                    name: newNodeName,
                    attributes: {identity: newNodeId},
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
        attributes: {identity: data[0].end.identity},
        children: []
    }
    
    //parsing tree
    for (let i = 0; i < data.length; i++) {
        const queue: RawNodeDatum[] = [];
        queue.unshift(tree as RawNodeDatum);
        while (queue.length > 0) {
            const curNode = queue.pop();
            if (curNode && curNode.children && curNode.attributes) {
                if (data[i].end.identity === curNode.attributes.identity) {
                    curNode.children.push({
                        name: data[i].start.properties.FirstName,
                        attributes: {
                            identity: data[i].start.identity
                        },
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