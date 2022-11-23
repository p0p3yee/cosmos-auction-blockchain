import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateAuction } from "./types/auction/auction/tx";
import { MsgPlaceBid } from "./types/auction/auction/tx";
import { MsgFinalizeAuction } from "./types/auction/auction/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/auction.auction.MsgCreateAuction", MsgCreateAuction],
    ["/auction.auction.MsgPlaceBid", MsgPlaceBid],
    ["/auction.auction.MsgFinalizeAuction", MsgFinalizeAuction],
    
];

export { msgTypes }