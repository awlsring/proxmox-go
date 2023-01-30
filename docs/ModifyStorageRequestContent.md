# ModifyStorageRequestContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocksize** | Pointer to **string** |  | [optional] 
**Bwlimit** | Pointer to **string** |  | [optional] 
**ComstarHg** | Pointer to **string** |  | [optional] 
**ComstarTg** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** | Allowed content types. | [optional] 
**DataPool** | Pointer to **string** | Data pool, only for erasure coding | [optional] 
**Delete** | Pointer to **string** | A list of settings to delete. | [optional] 
**Digest** | Pointer to **string** | Prevent change if current configuration has a different SHA1 digest. | [optional] 
**Disable** | Pointer to **bool** | Disable the storage | [optional] 
**Domain** | Pointer to **string** | CIFS domain | [optional] 
**EncryptionKey** | Pointer to **string** | Encryption key for storage. Use &#39;autogen&#39; to generate one automatically without passphrase. | [optional] 
**Fingerprint** | Pointer to **string** | Certificate SHA256 fingerprint | [optional] 
**Format** | Pointer to **string** | Default image format | [optional] 
**FsName** | Pointer to **string** | The ceph filesystem name | [optional] 
**Fuse** | Pointer to **bool** | Mount CephFS through FUSE | [optional] 
**IsMountpoint** | Pointer to **string** | Assume the given path is an externally managed mountpoint. This is a string boolean, using &#39;yes&#39; or &#39;no&#39; | [optional] 
**Keyring** | Pointer to **string** | Client keyring contents | [optional] 
**Krbd** | Pointer to **bool** | Always access rbd through krbd kernel module | [optional] 
**LioTpg** | Pointer to **string** | Target portal group for Linux LIO targets | [optional] 
**MasterPubkey** | Pointer to **string** | Base64 encoded, PEM formatted public RSA key | [optional] 
**MaxProtectedBackups** | Pointer to **float32** | Mac number of protected backups. Defaults to unlimited | [optional] 
**Maxfiles** | Pointer to **float32** | Maximum number of files. | [optional] 
**Mkdir** | Pointer to **bool** | Create the directory if doesnt exist. Defaults to true. | [optional] 
**Monhost** | Pointer to **string** | IP addresses of monitors. | [optional] 
**Mountpoint** | Pointer to **string** | The mount point | [optional] 
**Namespace** | Pointer to **string** | The namespace | [optional] 
**Nocow** | Pointer to **float32** | Set the NOCOW flag on files. This is a boolean integer, using 0 for false or 1 for true. | [optional] 
**Nodes** | Pointer to **string** | Comma seperated list of node names | [optional] 
**Nowritecache** | Pointer to **bool** | Disable write caching on the target. | [optional] 
**Options** | Pointer to **string** | NFS mount options. | [optional] 
**Password** | Pointer to **string** | Password for accessing the datastore | [optional] 
**Pool** | Pointer to **string** | The pool name | [optional] 
**Port** | Pointer to **float32** | For non default port | [optional] 
**Preallocation** | Pointer to [**StoragePreallocation**](StoragePreallocation.md) |  | [optional] 
**PruneBackups** | Pointer to **string** | The retention options. | [optional] 
**Saferemove** | Pointer to **bool** | Zero out data when removing LVs | [optional] 
**SaferemoveThroughput** | Pointer to **string** | Wipe throughput. | [optional] 
**Server** | Pointer to **string** | Server IP or DNS name | [optional] 
**Server2** | Pointer to **string** | Backup server IP or DNS name | [optional] 
**Shared** | Pointer to **bool** | Indicates if the storage is shared. | [optional] 
**Smbversion** | Pointer to [**StorageSMBVersion**](StorageSMBVersion.md) |  | [optional] 
**Sparse** | Pointer to **bool** | Use sparse volumes. | [optional] 
**Subdir** | Pointer to **string** | The subdir to mount. | [optional] 
**TaggedOnly** | Pointer to **bool** | Only use logical volumes tagged with &#39;pve-vm-ID&#39;. | [optional] 
**Transport** | Pointer to [**StorageTransport**](StorageTransport.md) |  | [optional] 
**Username** | Pointer to **string** | The username to use. | [optional] 

## Methods

### NewModifyStorageRequestContent

`func NewModifyStorageRequestContent() *ModifyStorageRequestContent`

NewModifyStorageRequestContent instantiates a new ModifyStorageRequestContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyStorageRequestContentWithDefaults

`func NewModifyStorageRequestContentWithDefaults() *ModifyStorageRequestContent`

NewModifyStorageRequestContentWithDefaults instantiates a new ModifyStorageRequestContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocksize

`func (o *ModifyStorageRequestContent) GetBlocksize() string`

GetBlocksize returns the Blocksize field if non-nil, zero value otherwise.

### GetBlocksizeOk

`func (o *ModifyStorageRequestContent) GetBlocksizeOk() (*string, bool)`

GetBlocksizeOk returns a tuple with the Blocksize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocksize

`func (o *ModifyStorageRequestContent) SetBlocksize(v string)`

SetBlocksize sets Blocksize field to given value.

### HasBlocksize

`func (o *ModifyStorageRequestContent) HasBlocksize() bool`

HasBlocksize returns a boolean if a field has been set.

### GetBwlimit

`func (o *ModifyStorageRequestContent) GetBwlimit() string`

GetBwlimit returns the Bwlimit field if non-nil, zero value otherwise.

### GetBwlimitOk

`func (o *ModifyStorageRequestContent) GetBwlimitOk() (*string, bool)`

GetBwlimitOk returns a tuple with the Bwlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwlimit

`func (o *ModifyStorageRequestContent) SetBwlimit(v string)`

SetBwlimit sets Bwlimit field to given value.

### HasBwlimit

`func (o *ModifyStorageRequestContent) HasBwlimit() bool`

HasBwlimit returns a boolean if a field has been set.

### GetComstarHg

`func (o *ModifyStorageRequestContent) GetComstarHg() string`

GetComstarHg returns the ComstarHg field if non-nil, zero value otherwise.

### GetComstarHgOk

`func (o *ModifyStorageRequestContent) GetComstarHgOk() (*string, bool)`

GetComstarHgOk returns a tuple with the ComstarHg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComstarHg

`func (o *ModifyStorageRequestContent) SetComstarHg(v string)`

SetComstarHg sets ComstarHg field to given value.

### HasComstarHg

`func (o *ModifyStorageRequestContent) HasComstarHg() bool`

HasComstarHg returns a boolean if a field has been set.

### GetComstarTg

`func (o *ModifyStorageRequestContent) GetComstarTg() string`

GetComstarTg returns the ComstarTg field if non-nil, zero value otherwise.

### GetComstarTgOk

`func (o *ModifyStorageRequestContent) GetComstarTgOk() (*string, bool)`

GetComstarTgOk returns a tuple with the ComstarTg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComstarTg

`func (o *ModifyStorageRequestContent) SetComstarTg(v string)`

SetComstarTg sets ComstarTg field to given value.

### HasComstarTg

`func (o *ModifyStorageRequestContent) HasComstarTg() bool`

HasComstarTg returns a boolean if a field has been set.

### GetContent

`func (o *ModifyStorageRequestContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModifyStorageRequestContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModifyStorageRequestContent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ModifyStorageRequestContent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDataPool

`func (o *ModifyStorageRequestContent) GetDataPool() string`

GetDataPool returns the DataPool field if non-nil, zero value otherwise.

### GetDataPoolOk

`func (o *ModifyStorageRequestContent) GetDataPoolOk() (*string, bool)`

GetDataPoolOk returns a tuple with the DataPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPool

`func (o *ModifyStorageRequestContent) SetDataPool(v string)`

SetDataPool sets DataPool field to given value.

### HasDataPool

`func (o *ModifyStorageRequestContent) HasDataPool() bool`

HasDataPool returns a boolean if a field has been set.

### GetDelete

`func (o *ModifyStorageRequestContent) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ModifyStorageRequestContent) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ModifyStorageRequestContent) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ModifyStorageRequestContent) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDigest

`func (o *ModifyStorageRequestContent) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *ModifyStorageRequestContent) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *ModifyStorageRequestContent) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *ModifyStorageRequestContent) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDisable

`func (o *ModifyStorageRequestContent) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *ModifyStorageRequestContent) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *ModifyStorageRequestContent) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *ModifyStorageRequestContent) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDomain

`func (o *ModifyStorageRequestContent) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ModifyStorageRequestContent) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ModifyStorageRequestContent) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ModifyStorageRequestContent) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *ModifyStorageRequestContent) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *ModifyStorageRequestContent) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *ModifyStorageRequestContent) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *ModifyStorageRequestContent) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetFingerprint

`func (o *ModifyStorageRequestContent) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ModifyStorageRequestContent) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ModifyStorageRequestContent) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ModifyStorageRequestContent) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFormat

`func (o *ModifyStorageRequestContent) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ModifyStorageRequestContent) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ModifyStorageRequestContent) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ModifyStorageRequestContent) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFsName

`func (o *ModifyStorageRequestContent) GetFsName() string`

GetFsName returns the FsName field if non-nil, zero value otherwise.

### GetFsNameOk

`func (o *ModifyStorageRequestContent) GetFsNameOk() (*string, bool)`

GetFsNameOk returns a tuple with the FsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsName

`func (o *ModifyStorageRequestContent) SetFsName(v string)`

SetFsName sets FsName field to given value.

### HasFsName

`func (o *ModifyStorageRequestContent) HasFsName() bool`

HasFsName returns a boolean if a field has been set.

### GetFuse

`func (o *ModifyStorageRequestContent) GetFuse() bool`

GetFuse returns the Fuse field if non-nil, zero value otherwise.

### GetFuseOk

`func (o *ModifyStorageRequestContent) GetFuseOk() (*bool, bool)`

GetFuseOk returns a tuple with the Fuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuse

`func (o *ModifyStorageRequestContent) SetFuse(v bool)`

SetFuse sets Fuse field to given value.

### HasFuse

`func (o *ModifyStorageRequestContent) HasFuse() bool`

HasFuse returns a boolean if a field has been set.

### GetIsMountpoint

`func (o *ModifyStorageRequestContent) GetIsMountpoint() string`

GetIsMountpoint returns the IsMountpoint field if non-nil, zero value otherwise.

### GetIsMountpointOk

`func (o *ModifyStorageRequestContent) GetIsMountpointOk() (*string, bool)`

GetIsMountpointOk returns a tuple with the IsMountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMountpoint

`func (o *ModifyStorageRequestContent) SetIsMountpoint(v string)`

SetIsMountpoint sets IsMountpoint field to given value.

### HasIsMountpoint

`func (o *ModifyStorageRequestContent) HasIsMountpoint() bool`

HasIsMountpoint returns a boolean if a field has been set.

### GetKeyring

`func (o *ModifyStorageRequestContent) GetKeyring() string`

GetKeyring returns the Keyring field if non-nil, zero value otherwise.

### GetKeyringOk

`func (o *ModifyStorageRequestContent) GetKeyringOk() (*string, bool)`

GetKeyringOk returns a tuple with the Keyring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyring

`func (o *ModifyStorageRequestContent) SetKeyring(v string)`

SetKeyring sets Keyring field to given value.

### HasKeyring

`func (o *ModifyStorageRequestContent) HasKeyring() bool`

HasKeyring returns a boolean if a field has been set.

### GetKrbd

`func (o *ModifyStorageRequestContent) GetKrbd() bool`

GetKrbd returns the Krbd field if non-nil, zero value otherwise.

### GetKrbdOk

`func (o *ModifyStorageRequestContent) GetKrbdOk() (*bool, bool)`

GetKrbdOk returns a tuple with the Krbd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrbd

`func (o *ModifyStorageRequestContent) SetKrbd(v bool)`

SetKrbd sets Krbd field to given value.

### HasKrbd

`func (o *ModifyStorageRequestContent) HasKrbd() bool`

HasKrbd returns a boolean if a field has been set.

### GetLioTpg

`func (o *ModifyStorageRequestContent) GetLioTpg() string`

GetLioTpg returns the LioTpg field if non-nil, zero value otherwise.

### GetLioTpgOk

`func (o *ModifyStorageRequestContent) GetLioTpgOk() (*string, bool)`

GetLioTpgOk returns a tuple with the LioTpg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLioTpg

`func (o *ModifyStorageRequestContent) SetLioTpg(v string)`

SetLioTpg sets LioTpg field to given value.

### HasLioTpg

`func (o *ModifyStorageRequestContent) HasLioTpg() bool`

HasLioTpg returns a boolean if a field has been set.

### GetMasterPubkey

`func (o *ModifyStorageRequestContent) GetMasterPubkey() string`

GetMasterPubkey returns the MasterPubkey field if non-nil, zero value otherwise.

### GetMasterPubkeyOk

`func (o *ModifyStorageRequestContent) GetMasterPubkeyOk() (*string, bool)`

GetMasterPubkeyOk returns a tuple with the MasterPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPubkey

`func (o *ModifyStorageRequestContent) SetMasterPubkey(v string)`

SetMasterPubkey sets MasterPubkey field to given value.

### HasMasterPubkey

`func (o *ModifyStorageRequestContent) HasMasterPubkey() bool`

HasMasterPubkey returns a boolean if a field has been set.

### GetMaxProtectedBackups

`func (o *ModifyStorageRequestContent) GetMaxProtectedBackups() float32`

GetMaxProtectedBackups returns the MaxProtectedBackups field if non-nil, zero value otherwise.

### GetMaxProtectedBackupsOk

`func (o *ModifyStorageRequestContent) GetMaxProtectedBackupsOk() (*float32, bool)`

GetMaxProtectedBackupsOk returns a tuple with the MaxProtectedBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProtectedBackups

`func (o *ModifyStorageRequestContent) SetMaxProtectedBackups(v float32)`

SetMaxProtectedBackups sets MaxProtectedBackups field to given value.

### HasMaxProtectedBackups

`func (o *ModifyStorageRequestContent) HasMaxProtectedBackups() bool`

HasMaxProtectedBackups returns a boolean if a field has been set.

### GetMaxfiles

`func (o *ModifyStorageRequestContent) GetMaxfiles() float32`

GetMaxfiles returns the Maxfiles field if non-nil, zero value otherwise.

### GetMaxfilesOk

`func (o *ModifyStorageRequestContent) GetMaxfilesOk() (*float32, bool)`

GetMaxfilesOk returns a tuple with the Maxfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxfiles

`func (o *ModifyStorageRequestContent) SetMaxfiles(v float32)`

SetMaxfiles sets Maxfiles field to given value.

### HasMaxfiles

`func (o *ModifyStorageRequestContent) HasMaxfiles() bool`

HasMaxfiles returns a boolean if a field has been set.

### GetMkdir

`func (o *ModifyStorageRequestContent) GetMkdir() bool`

GetMkdir returns the Mkdir field if non-nil, zero value otherwise.

### GetMkdirOk

`func (o *ModifyStorageRequestContent) GetMkdirOk() (*bool, bool)`

GetMkdirOk returns a tuple with the Mkdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMkdir

`func (o *ModifyStorageRequestContent) SetMkdir(v bool)`

SetMkdir sets Mkdir field to given value.

### HasMkdir

`func (o *ModifyStorageRequestContent) HasMkdir() bool`

HasMkdir returns a boolean if a field has been set.

### GetMonhost

`func (o *ModifyStorageRequestContent) GetMonhost() string`

GetMonhost returns the Monhost field if non-nil, zero value otherwise.

### GetMonhostOk

`func (o *ModifyStorageRequestContent) GetMonhostOk() (*string, bool)`

GetMonhostOk returns a tuple with the Monhost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonhost

`func (o *ModifyStorageRequestContent) SetMonhost(v string)`

SetMonhost sets Monhost field to given value.

### HasMonhost

`func (o *ModifyStorageRequestContent) HasMonhost() bool`

HasMonhost returns a boolean if a field has been set.

### GetMountpoint

`func (o *ModifyStorageRequestContent) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *ModifyStorageRequestContent) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *ModifyStorageRequestContent) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *ModifyStorageRequestContent) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetNamespace

`func (o *ModifyStorageRequestContent) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ModifyStorageRequestContent) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ModifyStorageRequestContent) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ModifyStorageRequestContent) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNocow

`func (o *ModifyStorageRequestContent) GetNocow() float32`

GetNocow returns the Nocow field if non-nil, zero value otherwise.

### GetNocowOk

`func (o *ModifyStorageRequestContent) GetNocowOk() (*float32, bool)`

GetNocowOk returns a tuple with the Nocow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocow

`func (o *ModifyStorageRequestContent) SetNocow(v float32)`

SetNocow sets Nocow field to given value.

### HasNocow

`func (o *ModifyStorageRequestContent) HasNocow() bool`

HasNocow returns a boolean if a field has been set.

### GetNodes

`func (o *ModifyStorageRequestContent) GetNodes() string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ModifyStorageRequestContent) GetNodesOk() (*string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ModifyStorageRequestContent) SetNodes(v string)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ModifyStorageRequestContent) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNowritecache

`func (o *ModifyStorageRequestContent) GetNowritecache() bool`

GetNowritecache returns the Nowritecache field if non-nil, zero value otherwise.

### GetNowritecacheOk

`func (o *ModifyStorageRequestContent) GetNowritecacheOk() (*bool, bool)`

GetNowritecacheOk returns a tuple with the Nowritecache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowritecache

`func (o *ModifyStorageRequestContent) SetNowritecache(v bool)`

SetNowritecache sets Nowritecache field to given value.

### HasNowritecache

`func (o *ModifyStorageRequestContent) HasNowritecache() bool`

HasNowritecache returns a boolean if a field has been set.

### GetOptions

`func (o *ModifyStorageRequestContent) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ModifyStorageRequestContent) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ModifyStorageRequestContent) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ModifyStorageRequestContent) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPassword

`func (o *ModifyStorageRequestContent) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModifyStorageRequestContent) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModifyStorageRequestContent) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModifyStorageRequestContent) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPool

`func (o *ModifyStorageRequestContent) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ModifyStorageRequestContent) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ModifyStorageRequestContent) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ModifyStorageRequestContent) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPort

`func (o *ModifyStorageRequestContent) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModifyStorageRequestContent) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModifyStorageRequestContent) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ModifyStorageRequestContent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPreallocation

`func (o *ModifyStorageRequestContent) GetPreallocation() StoragePreallocation`

GetPreallocation returns the Preallocation field if non-nil, zero value otherwise.

### GetPreallocationOk

`func (o *ModifyStorageRequestContent) GetPreallocationOk() (*StoragePreallocation, bool)`

GetPreallocationOk returns a tuple with the Preallocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreallocation

`func (o *ModifyStorageRequestContent) SetPreallocation(v StoragePreallocation)`

SetPreallocation sets Preallocation field to given value.

### HasPreallocation

`func (o *ModifyStorageRequestContent) HasPreallocation() bool`

HasPreallocation returns a boolean if a field has been set.

### GetPruneBackups

`func (o *ModifyStorageRequestContent) GetPruneBackups() string`

GetPruneBackups returns the PruneBackups field if non-nil, zero value otherwise.

### GetPruneBackupsOk

`func (o *ModifyStorageRequestContent) GetPruneBackupsOk() (*string, bool)`

GetPruneBackupsOk returns a tuple with the PruneBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneBackups

`func (o *ModifyStorageRequestContent) SetPruneBackups(v string)`

SetPruneBackups sets PruneBackups field to given value.

### HasPruneBackups

`func (o *ModifyStorageRequestContent) HasPruneBackups() bool`

HasPruneBackups returns a boolean if a field has been set.

### GetSaferemove

`func (o *ModifyStorageRequestContent) GetSaferemove() bool`

GetSaferemove returns the Saferemove field if non-nil, zero value otherwise.

### GetSaferemoveOk

`func (o *ModifyStorageRequestContent) GetSaferemoveOk() (*bool, bool)`

GetSaferemoveOk returns a tuple with the Saferemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaferemove

`func (o *ModifyStorageRequestContent) SetSaferemove(v bool)`

SetSaferemove sets Saferemove field to given value.

### HasSaferemove

`func (o *ModifyStorageRequestContent) HasSaferemove() bool`

HasSaferemove returns a boolean if a field has been set.

### GetSaferemoveThroughput

`func (o *ModifyStorageRequestContent) GetSaferemoveThroughput() string`

GetSaferemoveThroughput returns the SaferemoveThroughput field if non-nil, zero value otherwise.

### GetSaferemoveThroughputOk

`func (o *ModifyStorageRequestContent) GetSaferemoveThroughputOk() (*string, bool)`

GetSaferemoveThroughputOk returns a tuple with the SaferemoveThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaferemoveThroughput

`func (o *ModifyStorageRequestContent) SetSaferemoveThroughput(v string)`

SetSaferemoveThroughput sets SaferemoveThroughput field to given value.

### HasSaferemoveThroughput

`func (o *ModifyStorageRequestContent) HasSaferemoveThroughput() bool`

HasSaferemoveThroughput returns a boolean if a field has been set.

### GetServer

`func (o *ModifyStorageRequestContent) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ModifyStorageRequestContent) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ModifyStorageRequestContent) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *ModifyStorageRequestContent) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetServer2

`func (o *ModifyStorageRequestContent) GetServer2() string`

GetServer2 returns the Server2 field if non-nil, zero value otherwise.

### GetServer2Ok

`func (o *ModifyStorageRequestContent) GetServer2Ok() (*string, bool)`

GetServer2Ok returns a tuple with the Server2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer2

`func (o *ModifyStorageRequestContent) SetServer2(v string)`

SetServer2 sets Server2 field to given value.

### HasServer2

`func (o *ModifyStorageRequestContent) HasServer2() bool`

HasServer2 returns a boolean if a field has been set.

### GetShared

`func (o *ModifyStorageRequestContent) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *ModifyStorageRequestContent) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *ModifyStorageRequestContent) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *ModifyStorageRequestContent) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetSmbversion

`func (o *ModifyStorageRequestContent) GetSmbversion() StorageSMBVersion`

GetSmbversion returns the Smbversion field if non-nil, zero value otherwise.

### GetSmbversionOk

`func (o *ModifyStorageRequestContent) GetSmbversionOk() (*StorageSMBVersion, bool)`

GetSmbversionOk returns a tuple with the Smbversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbversion

`func (o *ModifyStorageRequestContent) SetSmbversion(v StorageSMBVersion)`

SetSmbversion sets Smbversion field to given value.

### HasSmbversion

`func (o *ModifyStorageRequestContent) HasSmbversion() bool`

HasSmbversion returns a boolean if a field has been set.

### GetSparse

`func (o *ModifyStorageRequestContent) GetSparse() bool`

GetSparse returns the Sparse field if non-nil, zero value otherwise.

### GetSparseOk

`func (o *ModifyStorageRequestContent) GetSparseOk() (*bool, bool)`

GetSparseOk returns a tuple with the Sparse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparse

`func (o *ModifyStorageRequestContent) SetSparse(v bool)`

SetSparse sets Sparse field to given value.

### HasSparse

`func (o *ModifyStorageRequestContent) HasSparse() bool`

HasSparse returns a boolean if a field has been set.

### GetSubdir

`func (o *ModifyStorageRequestContent) GetSubdir() string`

GetSubdir returns the Subdir field if non-nil, zero value otherwise.

### GetSubdirOk

`func (o *ModifyStorageRequestContent) GetSubdirOk() (*string, bool)`

GetSubdirOk returns a tuple with the Subdir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdir

`func (o *ModifyStorageRequestContent) SetSubdir(v string)`

SetSubdir sets Subdir field to given value.

### HasSubdir

`func (o *ModifyStorageRequestContent) HasSubdir() bool`

HasSubdir returns a boolean if a field has been set.

### GetTaggedOnly

`func (o *ModifyStorageRequestContent) GetTaggedOnly() bool`

GetTaggedOnly returns the TaggedOnly field if non-nil, zero value otherwise.

### GetTaggedOnlyOk

`func (o *ModifyStorageRequestContent) GetTaggedOnlyOk() (*bool, bool)`

GetTaggedOnlyOk returns a tuple with the TaggedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedOnly

`func (o *ModifyStorageRequestContent) SetTaggedOnly(v bool)`

SetTaggedOnly sets TaggedOnly field to given value.

### HasTaggedOnly

`func (o *ModifyStorageRequestContent) HasTaggedOnly() bool`

HasTaggedOnly returns a boolean if a field has been set.

### GetTransport

`func (o *ModifyStorageRequestContent) GetTransport() StorageTransport`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ModifyStorageRequestContent) GetTransportOk() (*StorageTransport, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ModifyStorageRequestContent) SetTransport(v StorageTransport)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ModifyStorageRequestContent) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUsername

`func (o *ModifyStorageRequestContent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModifyStorageRequestContent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModifyStorageRequestContent) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModifyStorageRequestContent) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


