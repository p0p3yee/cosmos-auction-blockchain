// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgFinalizeAuction } from "./types/auction/auction/tx";
import { MsgCreateAuction } from "./types/auction/auction/tx";
import { MsgPlaceBid } from "./types/auction/auction/tx";


export { MsgFinalizeAuction, MsgCreateAuction, MsgPlaceBid };

type sendMsgFinalizeAuctionParams = {
  value: MsgFinalizeAuction,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateAuctionParams = {
  value: MsgCreateAuction,
  fee?: StdFee,
  memo?: string
};

type sendMsgPlaceBidParams = {
  value: MsgPlaceBid,
  fee?: StdFee,
  memo?: string
};


type msgFinalizeAuctionParams = {
  value: MsgFinalizeAuction,
};

type msgCreateAuctionParams = {
  value: MsgCreateAuction,
};

type msgPlaceBidParams = {
  value: MsgPlaceBid,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgFinalizeAuction({ value, fee, memo }: sendMsgFinalizeAuctionParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgFinalizeAuction: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgFinalizeAuction({ value: MsgFinalizeAuction.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgFinalizeAuction: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateAuction({ value, fee, memo }: sendMsgCreateAuctionParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateAuction: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateAuction({ value: MsgCreateAuction.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateAuction: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgPlaceBid({ value, fee, memo }: sendMsgPlaceBidParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgPlaceBid: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgPlaceBid({ value: MsgPlaceBid.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgPlaceBid: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgFinalizeAuction({ value }: msgFinalizeAuctionParams): EncodeObject {
			try {
				return { typeUrl: "/auction.auction.MsgFinalizeAuction", value: MsgFinalizeAuction.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgFinalizeAuction: Could not create message: ' + e.message)
			}
		},
		
		msgCreateAuction({ value }: msgCreateAuctionParams): EncodeObject {
			try {
				return { typeUrl: "/auction.auction.MsgCreateAuction", value: MsgCreateAuction.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateAuction: Could not create message: ' + e.message)
			}
		},
		
		msgPlaceBid({ value }: msgPlaceBidParams): EncodeObject {
			try {
				return { typeUrl: "/auction.auction.MsgPlaceBid", value: MsgPlaceBid.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgPlaceBid: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			AuctionAuction: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;