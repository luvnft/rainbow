// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateVolatility is the `updateVolatility` instruction.
type UpdateVolatility struct {
	Args *UpdateVolatilityArgs

	// [0] = [] state
	//
	// [1] = [WRITE] greeks
	//
	// [2] = [] zetaGroup
	//
	// [3] = [SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateVolatilityInstructionBuilder creates a new `UpdateVolatility` instruction builder.
func NewUpdateVolatilityInstructionBuilder() *UpdateVolatility {
	nd := &UpdateVolatility{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *UpdateVolatility) SetArgs(args UpdateVolatilityArgs) *UpdateVolatility {
	inst.Args = &args
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *UpdateVolatility) SetStateAccount(state ag_solanago.PublicKey) *UpdateVolatility {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateVolatility) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetGreeksAccount sets the "greeks" account.
func (inst *UpdateVolatility) SetGreeksAccount(greeks ag_solanago.PublicKey) *UpdateVolatility {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(greeks).WRITE()
	return inst
}

// GetGreeksAccount gets the "greeks" account.
func (inst *UpdateVolatility) GetGreeksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetZetaGroupAccount sets the "zetaGroup" account.
func (inst *UpdateVolatility) SetZetaGroupAccount(zetaGroup ag_solanago.PublicKey) *UpdateVolatility {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(zetaGroup)
	return inst
}

// GetZetaGroupAccount gets the "zetaGroup" account.
func (inst *UpdateVolatility) GetZetaGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateVolatility) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateVolatility {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateVolatility) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst UpdateVolatility) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateVolatility,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateVolatility) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateVolatility) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Greeks is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ZetaGroup is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *UpdateVolatility) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateVolatility")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   greeks", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("zetaGroup", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    admin", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj UpdateVolatility) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateVolatility) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateVolatilityInstruction declares a new UpdateVolatility instruction with the provided parameters and accounts.
func NewUpdateVolatilityInstruction(
	// Parameters:
	args UpdateVolatilityArgs,
	// Accounts:
	state ag_solanago.PublicKey,
	greeks ag_solanago.PublicKey,
	zetaGroup ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *UpdateVolatility {
	return NewUpdateVolatilityInstructionBuilder().
		SetArgs(args).
		SetStateAccount(state).
		SetGreeksAccount(greeks).
		SetZetaGroupAccount(zetaGroup).
		SetAdminAccount(admin)
}