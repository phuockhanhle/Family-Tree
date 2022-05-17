import { Box } from "@mui/material";
import { useState } from "react";
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

export const DashBoard = () => {

    const [tree, setTree] = useState<RawNodeDatum | RawNodeDatum[]>({
        name: "Root",
        children: []
    });
    const [node, setNode] = useState<undefined | TreeNodeDatum>(undefined);
    const close = () => setNode(undefined);
    const handleSubmit = (name: string) => {
        if (node) {
            const newTree = bfs(node.name, tree, name);
            if (newTree) {
                setTree(newTree);
            }
        }
    };

    return (
        <>
            <Box sx={{
                width: '100vw',
                height: '100vh'
            }}>
                <Tree data={tree} translate={{ x: 200, y: 300 }} onNodeClick={(datum) => { setNode(datum.data) }} />
                <AddFamilyModal
                    isOpen={Boolean(node)}
                    onClose={close}
                    onSubmit={handleSubmit} />
            </Box>
        </>


    );
}

function bfs(name: string, tree: RawNodeDatum | RawNodeDatum[], newNodeName: string) {
    const queue: RawNodeDatum[] = [];
    queue.unshift(tree as RawNodeDatum);

    while (queue.length > 0) {
        const curNode = queue.pop();
        console.log(curNode);
        if (curNode && curNode.children) {
            if (curNode.name === name) {
                curNode.children.push({
                    name: newNodeName,
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