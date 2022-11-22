import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgPlaceBid } from "./types/auction/auction/tx";
import { MsgCreateAuction } from "./types/auction/auction/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/auction.auction.MsgPlaceBid", MsgPlaceBid],
    ["/auction.auction.MsgCreateAuction", MsgCreateAuction],
    
];

export { msgTypes }